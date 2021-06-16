package mvc

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"

	icore "github.com/ipfs/interface-go-ipfs-core"
)

func ChatCreateRecord(icapi icore.CoreAPI, db *Sql, value string) (vo.ChatRecordParams, error) {

	var msg vo.ChatRecordParams
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

	msg.Id = genRecordID(msg.FromId, msg.ToId)
	if msg.Ptime == 0 {
		msg.Ptime = time.Now().Unix()
	}

	// 检查是否存在
	var count int64
	err = db.DB.QueryRow("SELECT count(id) FROM chat_record WHERE id = ?", msg.Id).Scan(&count)
	if err != nil {
		sugar.Log.Error("Query chat_record failed.Err is", err)
		return msg, err
	}

	if count == 0 {
		res, err := db.DB.Exec("INSERT INTO chat_record (id, name, img, from_id, to_id, ptime, last_msg) VALUES (?, ?, ?, ?, ?, ?, ?)", msg.Id, msg.Name, msg.Img, msg.FromId, msg.ToId, msg.Ptime, msg.LastMsg)
		if err != nil {
			sugar.Log.Error("INSERT INTO chat_record is Failed.", err)
			return msg, err
		}

		_, err = res.LastInsertId()
		if err != nil {
			sugar.Log.Error("INSERT INTO chat_record is Failed2.", err)
			return msg, err
		}
	}

	topic := vo.MSG_LISTEN_PREFIX + msg.ToId

	sugar.Log.Info("publish topic: ", topic)

	msgBytes, err := json.Marshal(map[string]interface{}{
		"type": vo.MSG_TYPE_RECORD,
		"data": msg,
	})
	if err != nil {
		sugar.Log.Error("Message record marshal failed.Err is", err)
		return msg, err
	}

	sugar.Log.Info("publish data: ", string(msgBytes))

	ctx := context.Background()

	err = icapi.PubSub().Publish(ctx, topic, msgBytes)
	if err != nil {
		sugar.Log.Error("publish failed.", err)
		return msg, err
	}

	sugar.Log.Info("publish success")

	return msg, nil
}

func genRecordID(fromID, toID string) string {
	if fromID < toID {
		return toID + "_" + fromID
	} else {
		return fromID + "_" + toID
	}
}
