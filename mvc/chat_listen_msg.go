package mvc

import (
	"context"
	bsql "database/sql"
	"encoding/json"
	"errors"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"

	ipfsCore "github.com/ipfs/go-ipfs/core"
)

var ErrorAffectZero error = errors.New("row affect zero")
var ErrorRowNotExists error = errors.New("row not exists")
var ErrorRowIsExists error = errors.New("row is exists")

func ChatListenMsg(ipfsNode *ipfsCore.IpfsNode, db *Sql, token string, clh vo.ChatListenHandler) error {

	//校验 token 是否 满足
	claim, b := jwt.JwtVeriyToken(token)
	if !b {
		return errors.New("token 失效")
	}
	sugar.Log.Info("claim := ", claim)
	userid := claim["UserId"].(string)

	topic := vo.MSG_LISTEN_PREFIX + userid
	sugar.Log.Info("subscrib topic: ", topic)

	go func(topic string) {
		ctx := context.Background()

		sub, err := ipfsNode.PubSub.Subscribe(topic)
		// defer sub.Close()

		if err != nil {
			sugar.Log.Error("subscribe failed.", err)
			return
		}

		for {
			data, err := sub.Next(ctx)
			if err != nil {
				sugar.Log.Error("subscribe failed.", err)
				break
			}

			var msg vo.ChatListenParams

			sugar.Log.Debugf("receive: %s\n", data.Data)

			err = json.Unmarshal(data.Data, &msg)
			if err != nil {
				sugar.Log.Error("data unmarshal failed", err)
				continue
			}

			if msg.Type == vo.MSG_TYPE_RECORD {

				var tmp vo.ChatRecordParams
				json1, _ := json.Marshal(msg.Data)
				json.Unmarshal(json1, &tmp)

				res, err := handleAddRecordMsg(db, tmp)
				if err != nil {
					sugar.Log.Error("handle add record failed.", err)
					continue
				}

				msg.Data = res
				jsonStr, _ := json.Marshal(msg)
				clh.HandlerChat(string(jsonStr))

			} else if msg.Type == vo.MSG_TYPE_NEW {

				var tmp vo.ChatMsgParams
				json1, _ := json.Marshal(msg.Data)
				json.Unmarshal(json1, &tmp)

				res, err := handleNewMsg(db, tmp)
				if err != nil {
					sugar.Log.Error("handle add message failed.", err)
					continue
				}
				msg.Data = res
				jsonStr, _ := json.Marshal(msg)
				clh.HandlerChat(string(jsonStr))

			} else if msg.Type == vo.MSG_TYPE_WITHDRAW {

				var tmp vo.ChatMsgParams
				json1, _ := json.Marshal(msg.Data)
				json.Unmarshal(json1, &tmp)

				res, err := handleWithdrawMsg(db, tmp)
				if err != nil {
					sugar.Log.Error("handle withdraw message failed.", err)
					continue
				}
				msg.Data = res
				jsonStr, _ := json.Marshal(msg)
				clh.HandlerChat(string(jsonStr))

			} else {
				sugar.Log.Error("unsupport msg type", err)
				continue
			}
		}
	}(topic)

	// 发布消息
	return nil
}

// handleAddRecordMsg 创建会话
func handleAddRecordMsg(db *Sql, msg vo.ChatRecordParams) (vo.ChatRecordParams, error) {

	var record vo.ChatRecordParams
	err := db.DB.QueryRow("SELECT id, name, img, from_id, to_id, ptime, last_msg FROM chat_record WHERE id = ?", msg.Id).Scan(&record.Id, &record.Name, &record.Img, &record.FromId, &record.ToId, &record.Ptime, &record.LastMsg)

	switch err {
	case bsql.ErrNoRows:
		res, err := db.DB.Exec("INSERT INTO chat_record (id, name, img, from_id, to_id, ptime, last_msg) values (?, ?, ?, ?, ?, ?, ?)",
			msg.Id, msg.Name, msg.Img, msg.FromId, msg.ToId, msg.Ptime, msg.LastMsg)
		if err != nil {
			return msg, err
		}
		num, err := res.LastInsertId()
		if err != nil {
			return msg, err
		} else if num == 0 {
			return msg, err
		}

		return msg, nil
	case nil:
		if err != nil {
			return msg, err
		}

		if record.Ptime > msg.Ptime && record.FromId != msg.FromId {
			res, err := db.DB.Exec("UPDATE chat_record SET from_id = ?, to_id = ?, ptime = ?, last_msg = ? WHERE id = ?", msg.FromId, msg.ToId, msg.Ptime, msg.LastMsg, msg.Id)
			if err != nil {
				return msg, err
			}
			num, err := res.RowsAffected()
			if err != nil {
				return msg, err
			} else if num == 0 {
				return msg, err
			}
		}
		return msg, nil

	default:
		return msg, err
	}

}

// handleWithdrawMsg 撤销消息
func handleWithdrawMsg(db *Sql, msg vo.ChatMsgParams) (vo.ChatMsgParams, error) {
	var cMsg vo.ChatMsgParams

	err := db.DB.QueryRow("SELECT id, content_type, content, from_id, to_id, ptime, is_with_draw, is_read, record_id FROM chat_msg WHERE id = ?", msg.Id).Scan(&cMsg.Id, &cMsg.ContentType, &cMsg.Content, &cMsg.FromId, &cMsg.ToId, &cMsg.Ptime, &cMsg.IsWithdraw, &cMsg.IsRead, &cMsg.RecordId)

	switch err {
	case bsql.ErrNoRows:
		return msg, err
	case nil:
		res, err := db.DB.Exec("UPDATE chat_msg SET is_with_draw = 1 WHERE id = ? and from_id = ?", msg.Id, msg.FromId)
		if err != nil {
			return msg, err
		}
		num, err := res.RowsAffected()
		if err != nil {
			return msg, err
		} else if num == 0 {
			return msg, ErrorAffectZero
		}

		cMsg.IsWithdraw = 1
		return cMsg, nil
	default:
		return msg, err
	}
}

// handleNewMsg 新增消息
func handleNewMsg(db *Sql, msg vo.ChatMsgParams) (vo.ChatMsgParams, error) {

	var cMsg vo.ChatMsgParams

	err := db.DB.QueryRow("SELECT id, content_type, content, from_id, to_id, ptime, is_with_draw, is_read, record_id FROM chat_msg WHERE id = ?", msg.Id).Scan(&cMsg.Id, &cMsg.ContentType, &cMsg.Content, &cMsg.FromId, &cMsg.ToId, &cMsg.Ptime, &cMsg.IsWithdraw, &cMsg.IsRead, &cMsg.RecordId)
	switch err {
	case bsql.ErrNoRows:
		res, err := db.DB.Exec("INSERT INTO chat_msg (id, content_type, content, from_id, to_id, ptime, is_with_draw, is_read, record_id) values (?, ?, ?, ?, ?, ?, ?, ?, ?)",
			msg.Id, msg.ContentType, msg.Content, msg.FromId, msg.ToId, msg.Ptime, msg.IsWithdraw, msg.IsRead, msg.RecordId)
		if err != nil {
			return msg, err
		}
		_, err = res.LastInsertId()
		if err != nil {
			return msg, err
		}

		_, err = db.DB.Exec("UPDATE chat_record SET last_msg = ? WHERE id = ?", msg.Content, msg.RecordId)
		if err != nil {
			return msg, err
		}

		return msg, nil

	case nil:
		return msg, ErrorRowIsExists
	default:
		return msg, err
	}
}
