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

				var tmp vo.ChatSwapRecordParams
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

				var tmp vo.ChatSwapMsgParams
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

				var tmp vo.ChatSwapWithdrawMsgParams
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
func handleAddRecordMsg(db *Sql, msg vo.ChatSwapRecordParams) (vo.ChatRecordInfo, error) {

	ret := vo.ChatRecordInfo{
		Id:      msg.Id,
		Name:    msg.Name,
		Img:     msg.Img,
		FromId:  msg.FromId,
		Toid:    msg.ToId,
		Ptime:   msg.Ptime,
		LastMsg: msg.LastMsg,

		UserName: "",
		Phone:    "",
		PeerId:   "",
		NickName: "",
		Sex:      0,
	}

	var ptime int64
	err := db.DB.QueryRow("SELECT ptime FROM chat_record WHERE id = ?", msg.Id).Scan(&ptime)

	switch err {
	case bsql.ErrNoRows:
		res, err := db.DB.Exec("INSERT INTO chat_record (id, name, from_id, to_id, ptime, last_msg) values (?, ?, ?, ?, ?, ?)",
			msg.Id, msg.Name, msg.FromId, msg.ToId, msg.Ptime, msg.LastMsg)
		if err != nil {
			return ret, err
		}
		_, err = res.LastInsertId()
		if err != nil {
			return ret, err
		}

		// 查询对方信息
		err = db.DB.QueryRow("SELECT peer_id, name, phone, sex, nickname, img FROM cloud_user WHERE id = ?", msg.FromId).Scan(&ret.PeerId, &ret.UserName, &ret.Phone, &ret.Sex, &ret.NickName, &ret.Img)
		if err != nil {
			sugar.Log.Error("Query Peer User Failed. Err:", err)
			return ret, err
		}

		return ret, nil
	case nil:
		if ptime > msg.Ptime {
			res, err := db.DB.Exec("UPDATE chat_record SET from_id, to_id, ptime = ?, last_msg = ? WHERE id = ?", msg.FromId, msg.ToId, msg.Ptime, msg.LastMsg, msg.Id)
			if err != nil {
				return ret, err
			}
			num, err := res.RowsAffected()
			if err != nil {
				return ret, err
			} else if num == 0 {
				return ret, err
			}
		}
		return ret, vo.ErrorRowIsExists

	default:
		return ret, err
	}

}

// handleWithdrawMsg 撤销消息
func handleWithdrawMsg(db *Sql, msg vo.ChatSwapWithdrawMsgParams) (ChatMsg, error) {

	ret := ChatMsg{
		Id:     msg.MsgId,
		FromId: msg.FromId,
		ToId:   msg.ToId,
	}

	err := db.DB.QueryRow("SELECT id, content_type, content, from_id, to_id, ptime, is_with_draw, is_read, record_id FROM chat_msg WHERE id = ?", ret.Id).Scan(&ret.Id, &ret.ContentType, &ret.Content, &ret.FromId, &ret.ToId, &ret.Ptime, &ret.IsWithdraw, &ret.IsRead, &ret.RecordId)

	switch err {
	case bsql.ErrNoRows:
		return ret, err
	case nil:
		res, err := db.DB.Exec("UPDATE chat_msg SET is_with_draw = 1 WHERE id = ? and from_id = ?", ret.Id, ret.FromId)
		if err != nil {
			return ret, err
		}
		num, err := res.RowsAffected()
		if err != nil {
			return ret, err
		} else if num == 0 {
			return ret, vo.ErrorAffectZero
		}

		ret.IsWithdraw = 1
		return ret, nil
	default:
		return ret, err
	}
}

// handleNewMsg 新增消息
func handleNewMsg(db *Sql, msg vo.ChatSwapMsgParams) (ChatMsg, error) {

	var recordId string

	ret := ChatMsg{
		Id:          msg.Id,
		ContentType: msg.ContentType,
		Content:     msg.Content,
		FromId:      msg.FromId,
		ToId:        msg.ToId,
		Ptime:       msg.Ptime,
		IsWithdraw:  msg.IsWithdraw,
		IsRead:      msg.IsRead,
		RecordId:    msg.RecordId,
	}

	// 检查房间是否存在
	err := db.DB.QueryRow("SELECT id FROM chat_record WHERE id = ?", ret.RecordId).Scan(&recordId)
	switch err {
	case bsql.ErrNoRows:
		res, err := db.DB.Exec("INSERT INTO chat_record (id, name, from_id, to_id, ptime, last_msg) values (?, ?, ?, ?, ?, ?)",
			ret.RecordId, "...", ret.FromId, ret.ToId, ret.Ptime, ret.Content)
		if err != nil {
			return ret, err
		}
		_, err = res.LastInsertId()
		if err != nil {
			return ret, err
		}
	case nil:
		// nothing
	default:
		return ret, err
	}

	// 检查消息是否重复
	err = db.DB.QueryRow("SELECT id, content_type, content, from_id, to_id, ptime, is_with_draw, is_read, record_id FROM chat_msg WHERE id = ?", ret.Id).Scan(&ret.Id, &ret.ContentType, &ret.Content, &ret.FromId, &ret.ToId, &ret.Ptime, &ret.IsWithdraw, &ret.IsRead, &ret.RecordId)
	switch err {
	case bsql.ErrNoRows:
		res, err := db.DB.Exec("INSERT INTO chat_msg (id, content_type, content, from_id, to_id, ptime, is_with_draw, is_read, record_id) values (?, ?, ?, ?, ?, ?, ?, ?, ?)",
			ret.Id, ret.ContentType, ret.Content, ret.FromId, ret.ToId, ret.Ptime, ret.IsWithdraw, ret.IsRead, ret.RecordId)
		if err != nil {
			return ret, err
		}
		_, err = res.LastInsertId()
		if err != nil {
			return ret, err
		}

		_, err = db.DB.Exec("UPDATE chat_record SET last_msg = ?, ptime = ? WHERE id = ?", ret.Content, ret.Ptime, ret.RecordId)
		if err != nil {
			return ret, err
		}

		return ret, nil

	case nil:
		return ret, vo.ErrorRowIsExists
	default:
		return ret, err
	}
}
