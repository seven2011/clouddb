package mvc

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"

	"github.com/cosmopolitann/clouddb/jwt"

	ipfsCore "github.com/ipfs/go-ipfs/core"
)

func ChatWithdrawMsg(ipfsNode *ipfsCore.IpfsNode, db *Sql, value string) error {

	var msg vo.ChatWithdrawMsgParams
	err := json.Unmarshal([]byte(value), &msg)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return err
	}
	sugar.Log.Info("Marshal data is  ", msg)

	//校验 token 是否 满足
	claim, b := jwt.JwtVeriyToken(msg.Token)
	if !b {
		return errors.New("token 失效")
	}
	sugar.Log.Info("claim := ", claim)
	userid := claim["UserId"].(string)

	if userid != msg.FromId {
		sugar.Log.Error("token is not msg.from_id")
		return errors.New("token is not msg.from_id")
	}

	msg.Token = ""

	// 查询会话是否存在
	var recordID string
	var toID string
	var isWithdraw int64

	err = db.DB.QueryRow("SELECT to_id, is_with_draw, record_id FROM chat_msg WHERE id = ? and from_id = ?", msg.MsgId, msg.FromId).Scan(&toID, &isWithdraw, &recordID)
	if err != nil {
		sugar.Log.Error("SELECT id FROM chat_msg is failed.", err)
		return err
	}

	if isWithdraw == 0 {
		res, err := db.DB.Exec("UPDATE chat_msg SET is_with_draw = 1 WHERE id = ? and from_id = ?", msg.MsgId, msg.FromId)
		if err != nil {
			sugar.Log.Error("UPDATE chat_msg is withdraw failed.", err)
			return err
		}

		num, err := res.RowsAffected()
		if err != nil {
			sugar.Log.Error("UPDATE chat_msg is withdraw failed2.", err)
			return err
		} else if num == 0 {
			sugar.Log.Error("UPDATE chat_msg is withdraw failed3.", err)
			return errors.New("UPDATE chat_msg is withdraw failed3")
		}
	}

	msgBytes, err := json.Marshal(map[string]interface{}{
		"type": vo.MSG_TYPE_WITHDRAW,
		"data": msg,
	})
	if err != nil {
		sugar.Log.Error("marshal send msg failed.", err)
		return err
	}

	ipfsTopic, ok := TopicJoin.Load(vo.CHAT_MSG_SWAP_TOPIC)
	if !ok {
		ipfsTopic, err = ipfsNode.PubSub.Join(vo.CHAT_MSG_SWAP_TOPIC)
		if err != nil {
			sugar.Log.Error("PubSub.Join .Err is", err)
			return err
		}

		TopicJoin.Store(vo.CHAT_MSG_SWAP_TOPIC, ipfsTopic)
	}

	ctx := context.Background()

	err = ipfsTopic.Publish(ctx, msgBytes)
	if err != nil {
		sugar.Log.Error("publish failed.", err)
		return err
	}

	sugar.Log.Info("publish success")

	// 发布消息
	return nil

}
