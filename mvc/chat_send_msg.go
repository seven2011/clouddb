package mvc

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/utils"
	"github.com/cosmopolitann/clouddb/vo"

	"github.com/cosmopolitann/clouddb/jwt"

	ipfsCore "github.com/ipfs/go-ipfs/core"
)

func ChatSendMsg(ipfsNode *ipfsCore.IpfsNode, db *Sql, value string) (vo.ChatMsgParams, error) {

	var msg vo.ChatMsgParams
	err := json.Unmarshal([]byte(value), &msg)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return msg, err
	}
	sugar.Log.Info("Marshal data is  ", msg)

	//校验 token 是否 满足
	claim, b := jwt.JwtVeriyToken(msg.Token)
	if !b {
		return msg, errors.New("token 失效")
	}
	sugar.Log.Info("claim := ", claim)
	userid := claim["UserId"].(string)

	if userid != msg.FromId {
		sugar.Log.Error("token is not msg.from_id")
		return msg, errors.New("token is not msg.from_id")
	}

	msg.Token = ""
	msg.Id = strconv.FormatInt(utils.SnowId(), 10)

	if len(msg.RecordId) == 0 {
		msg.RecordId = genRecordID(msg.FromId, msg.ToId)
	}

	if msg.Ptime == 0 {
		msg.Ptime = time.Now().Unix()
	}

	res, err := db.DB.Exec(
		"INSERT INTO chat_msg (id, content_type, content, from_id, to_id, ptime, is_with_draw, is_read, record_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		msg.Id, msg.ContentType, msg.Content, msg.FromId, msg.ToId, msg.Ptime, msg.IsWithdraw, msg.IsRead, msg.RecordId)
	if err != nil {
		sugar.Log.Error("INSERT INTO chat_msg is Failed.", err)
		return msg, err
	}

	_, err = res.LastInsertId()
	if err != nil {
		sugar.Log.Error("INSERT INTO chat_msg is Failed2.", err)
		return msg, err
	}

	msgBytes, err := json.Marshal(map[string]interface{}{
		"type": vo.MSG_TYPE_NEW,
		"data": msg,
	})
	if err != nil {
		sugar.Log.Error("marshal send msg failed.", err)
		return msg, err
	}

	ipfsTopic, ok := TopicJoin.Load(vo.CHAT_MSG_SWAP_TOPIC)
	if !ok {
		ipfsTopic, err = ipfsNode.PubSub.Join(vo.CHAT_MSG_SWAP_TOPIC)
		if err != nil {
			sugar.Log.Error("PubSub.Join .Err is", err)
			return msg, err
		}

		TopicJoin.Store(vo.CHAT_MSG_SWAP_TOPIC, ipfsTopic)
	}

	ctx := context.Background()

	err = ipfsTopic.Publish(ctx, msgBytes)
	if err != nil {
		sugar.Log.Error("publish failed.", err)
		return msg, err
	}

	sugar.Log.Info("publish success")

	// 发布消息
	return msg, nil

}
