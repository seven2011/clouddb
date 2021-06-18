package mvc

import (
	"context"
	bsql "database/sql"
	"encoding/json"
	"errors"
	"runtime/debug"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	pubsub "github.com/libp2p/go-libp2p-pubsub"

	ipfsCore "github.com/ipfs/go-ipfs/core"
)

func ChatListenMsg(ipfsNode *ipfsCore.IpfsNode, db *Sql, token string, clh vo.ChatListenHandler) error {

	//校验 token 是否 满足
	claim, b := jwt.JwtVeriyToken(token)
	if !b {
		return errors.New("token 失效")
	}
	sugar.Log.Info("claim := ", claim)
	userId := claim["UserId"].(string)

	var err error
	ctx := context.Background()

	ipfsTopic, ok := TopicJoin.Load(vo.CHAT_MSG_SWAP_TOPIC)
	if !ok {
		ipfsTopic, err = ipfsNode.PubSub.Join(vo.CHAT_MSG_SWAP_TOPIC)
		if err != nil {
			sugar.Log.Error("PubSub.Join .Err is", err)
			return err
		}

		TopicJoin.Store(vo.CHAT_MSG_SWAP_TOPIC, ipfsTopic)
	}

	go func(userId string, ipfsTopic *pubsub.Topic) {

		defer func() {
			// 错误捕获
			if r := recover(); r != nil {
				sugar.Log.Error("ChatListenMsg goroutine panic occure, err:", r)
				sugar.Log.Error("stack:", debug.Stack())
			}
		}()

		sub, err := ipfsTopic.Subscribe()
		if err != nil {
			sugar.Log.Error("subscribe failed.", err)
			return
		}

		var msg vo.ChatListenParams

		for {
			data, err := sub.Next(ctx)
			if err != nil {
				sugar.Log.Error("subscribe failed.", err)
				return
			}
			sugar.Log.Debugf("receive: %s\n", data.Data)

			msg = vo.ChatListenParams{}

			err = json.Unmarshal(data.Data, &msg)
			if err != nil {
				sugar.Log.Error("data unmarshal failed", err)
				continue
			}

			if msg.Type == vo.MSG_TYPE_RECORD {

				var tmp vo.ChatRecordParams
				json1, _ := json.Marshal(msg.Data)
				json.Unmarshal(json1, &tmp)

				if tmp.ToId != userId {
					// not me
					continue
				}

				res, err := handleAddRecordMsg(db, tmp)
				if err != nil {
					if err != vo.ErrorRowIsExists {
						sugar.Log.Error("handle add record failed.", err)
					}
					continue
				}

				msg.Data = res
				jsonStr, _ := json.Marshal(msg)
				clh.HandlerChat(string(jsonStr))

			} else if msg.Type == vo.MSG_TYPE_NEW {

				var tmp vo.ChatMsgParams
				json1, _ := json.Marshal(msg.Data)
				json.Unmarshal(json1, &tmp)

				if tmp.ToId != userId {
					// not me
					continue
				}

				res, err := handleNewMsg(db, tmp)
				if err != nil {
					if err != vo.ErrorRowIsExists {
						sugar.Log.Error("handle add message failed.", err)
					}
					continue
				}
				msg.Data = res
				jsonStr, _ := json.Marshal(msg)
				clh.HandlerChat(string(jsonStr))

			} else if msg.Type == vo.MSG_TYPE_WITHDRAW {

				var tmp vo.ChatMsgParams
				json1, _ := json.Marshal(msg.Data)
				json.Unmarshal(json1, &tmp)

				if tmp.ToId != userId {
					// not me
					continue
				}

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
	}(userId, ipfsTopic)

	return nil
}

// handleAddRecordMsg 创建会话
func handleAddRecordMsg(db *Sql, msg vo.ChatRecordParams) (vo.ChatRecordParams, error) {

	var record vo.ChatRecordParams
	err := db.DB.QueryRow("SELECT id, name, img, from_id, to_id, ptime, last_msg FROM chat_record WHERE id = ?", msg.Id).Scan(&record.Id, &record.Name, &record.Img, &record.FromId, &record.ToId, &record.Ptime, &record.LastMsg)

	switch err {
	case bsql.ErrNoRows:
		// swap from_id and to_id
		msg.FromId, msg.ToId = msg.ToId, msg.FromId
		res, err := db.DB.Exec("INSERT INTO chat_record (id, name, img, from_id, to_id, ptime, last_msg) values (?, ?, ?, ?, ?, ?, ?)",
			msg.Id, msg.Name, msg.Img, msg.FromId, msg.ToId, msg.Ptime, msg.LastMsg)
		if err != nil {
			return msg, err
		}
		_, err = res.LastInsertId()
		if err != nil {
			return msg, err
		}

		return msg, nil
	case nil:
		if record.Ptime > msg.Ptime {
			res, err := db.DB.Exec("UPDATE chat_record SET ptime = ?, last_msg = ? WHERE id = ?", msg.Ptime, msg.LastMsg, msg.Id)
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
		return msg, vo.ErrorRowIsExists

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
			return msg, vo.ErrorAffectZero
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
	var count int64

	err := db.DB.QueryRow("SELECT count(id) WHERE id = ?", msg.Id).Scan(&count)

	if err == bsql.ErrNoRows {
		// swap from_id and to_id
		res, err := db.DB.Exec("INSERT INTO chat_record (id, name, img, from_id, to_id, ptime, last_msg) values (?, ?, ?, ?, ?, ?, ?)",
			msg.RecordId, "", "", msg.ToId, msg.FromId, msg.Ptime, msg.Content)
		if err != nil {
			return msg, err
		}
		_, err = res.LastInsertId()
		if err != nil {
			return msg, err
		}
	}

	err = db.DB.QueryRow("SELECT id, content_type, content, from_id, to_id, ptime, is_with_draw, is_read, record_id FROM chat_msg WHERE id = ?", msg.Id).Scan(&cMsg.Id, &cMsg.ContentType, &cMsg.Content, &cMsg.FromId, &cMsg.ToId, &cMsg.Ptime, &cMsg.IsWithdraw, &cMsg.IsRead, &cMsg.RecordId)
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
		return msg, vo.ErrorRowIsExists
	default:
		return msg, err
	}
}
