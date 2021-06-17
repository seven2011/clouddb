package mvc

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/utils"
	ipfsCore "github.com/ipfs/go-ipfs/core"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"strconv"
	"time"
)

func AddUser(ipfsNode *ipfsCore.IpfsNode,db *Sql, value string) error {
	//user string ==> user struct
	//Add sys_user
	//create snow id

	var user SysUser
	err := json.Unmarshal([]byte(value), &user)
	if err != nil {

	}
	sugar.Log.Info("params ：= ", user)

	l, e := FindIsExistUser(db, user)
	if e != nil {
		sugar.Log.Error("FindIsExistUser info is Failed.")
	}
	// l > 0 user is exist.
	sugar.Log.Error("-----------1")

	if l > 0 {
		sugar.Log.Error("user is exist.")
		return errors.New("user is exist.")
	}

	//inExist insert into sys_user.

	sugar.Log.Info("-----------用户 信息 ", user)

	id := utils.SnowId()
	//create now time
	//t:=time.Now().Format("2006-01-02 15:04:05")
	t := time.Now().Unix()
	stmt, err := db.DB.Prepare("INSERT INTO sys_user values(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert data to sys_user is failed.")
		return err
	}
	sid := strconv.FormatInt(id, 10)
	res, err := stmt.Exec(sid, user.PeerId, user.Name, user.Phone, user.Sex, t, t, user.NickName, user.Img)
	if err != nil {
		sugar.Log.Error("Insert data to sys_user is failed.", res)
		return err
	}
	c, _ := res.RowsAffected()
	sugar.Log.Info("~~~~~   Insert into sys_user data is Successful ~~~~~~", c)
	//生成 token
	// 手机号
	//token,err:=jwt.GenerateToken(user.Phone,60)

	//=====
	// publish msg
	topic:="/db-online-sync"
	sugar.Log.Info("发布主题:","/db-online-sync")
	sugar.Log.Info("发布消息:",value)
	//判断是否弃用
	var tp *pubsub.Topic
	var ok bool
	ctx := context.Background()
	if tp,ok = Topicmp["/db-online-sync"];ok == false {
		tp, err = ipfsNode.PubSub.Join(topic)
		if err != nil {
			return err
		}
		Topicmp[topic] = tp

	}
	sugar.Log.Info("--- 开始 发布的消息 ---")

	sugar.Log.Info("发布的消息:", value)
//=====

	//================================

	//第一步
	var s3 UserAd
	s3.Type = "receiveUserRegister"
	s3.Data = user
	//

	jsonBytes, err := json.Marshal(s3)
	if err != nil {
		sugar.Log.Info("--- 开始 发布的消息 ---")
		return err
	}
	sugar.Log.Info("--- 解析后的数据 返回给 转接服务器 ---",string(jsonBytes))


	//============================


//====

	err = tp.Publish(ctx,jsonBytes)
	if err != nil {
		sugar.Log.Error("发布错误:", err)
		return err
	}
	sugar.Log.Error("---  发布的消息  完成  ---")

	return nil
}


type UserAd struct {
	Type string `json:"type"`

	Data SysUser `json:"data"`
}

func FindIsExistUser(db *Sql, user SysUser) (int64, error) {
	var s SysUser
	sugar.Log.Info("start sys_user is exist local user info.")
	sugar.Log.Info("user info is ", user.Phone)
	sugar.Log.Info("user info is ", user)

	rows, _ := db.DB.Query("SELECT * FROM sys_user where phone=?", user.Phone)
	for rows.Next() {
		err := rows.Scan(&s.Id, &s.PeerId, &s.Name, &s.Phone, &s.Sex, &s.Ptime, &s.Utime, &s.NickName, &s.Img)
		if err != nil {
			sugar.Log.Error(" query is failed. ", err)

			return 0, err
		}
		sugar.Log.Info(" user info is ", s)
	}
	//is exist
	sugar.Log.Info(" FindOne data is ", s.Id)

	if s.Id != "" {
		return 1, nil
	} else {
		return 0, nil
	}

}
