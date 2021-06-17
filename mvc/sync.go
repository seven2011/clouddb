package mvc

import (
	"encoding/json"
	"errors"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	ipfsCore "github.com/ipfs/go-ipfs/core"

	"context"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"time"
)

func SyncUser(db *Sql, value string) error {

	var user SysUser
	err := json.Unmarshal([]byte(value), &user)
	if err != nil {
		sugar.Log.Error("解析失败:", err)
		return err
	}
	sugar.Log.Info("params ：= ", user)
	//l,e:= FindIsExistUser(db,user)
	//if e!=nil{
	//	sugar.Log.Error("FindIsExistUser info is Failed.")
	//}
	//// l > 0 user is exist.
	//sugar.Log.Error("-----------1")
	//
	//if l>0{
	//	sugar.Log.Error("user is exist.")
	//	return errors.New("user is exist.")
	//}

	//inExist insert into sys_user.
	//	id := utils.SnowId()
	//create now time
	//t:=time.Now().Format("2006-01-02 15:04:05")
	t := time.Now().Unix()
	stmt, err := db.DB.Prepare("INSERT INTO sys_user values(?,?,?,?,?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert data to sys_user is failed.")
		return err
	}

	//sid := strconv.FormatInt(user.Id, 10)
	res, err := stmt.Exec(user.Id, user.PeerId, user.Name, user.Phone, user.Sex, t, t, user.NickName)
	if err != nil {
		sugar.Log.Error("Insert data to sys_user is failed.", res)
		return err
	}
	c, _ := res.RowsAffected()
	sugar.Log.Info("~~~~~   Sync into sys_user data is Successful ~~~~~~", c)
	//生成 token
	// 手机号
	//token,err:=jwt.GenerateToken(user.Phone,60)

	return nil
}

// 文章

func SyncArticle(db *Sql, value string) error {
	var art vo.ArticleAddParams
	err := json.Unmarshal([]byte(value), &art)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return errors.New("解析字段错误")
	}
	sugar.Log.Info("Marshal data is  ", art)
	//id := utils.SnowId()
	//t := time.Now().Format("2006-01-02 15:04:05")
	t := time.Now().Unix()

	stmt, err := db.DB.Prepare("INSERT INTO article values(?,?,?,?,?,?,?,?,?,?)")

	if err != nil {
		sugar.Log.Error("Insert into article table is failed.", err)
		return errors.New("插入article 表数据 失败")

	}
	//sid := strconv.FormatInt(id, 10)
	stmt.QueryRow()
	res, err := stmt.Exec(art.Id, art.UserId, art.Accesstory, art.AccesstoryType, art.Text, art.Tag, t, 0, art.Title, 0)
	if err != nil {
		sugar.Log.Error("Insert into article  is Failed.", err)
		return errors.New("插入数据失败")
	}
	sugar.Log.Info("Insert into article  is successful.")
	l, _ := res.RowsAffected()
	if l == 0 {
		return errors.New("插入数据失败")
	}
	return nil

}

//

func SyncAticlePlay(db *Sql, value string) error {
	//更新字段  is_ike = 1
	var art vo.SyncArticleGiveLikeParams
	err := json.Unmarshal([]byte(value), &art)

	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", art)
	//插入新的一条记录
	//id := utils.SnowId()
	stmt, err := db.DB.Prepare("INSERT INTO article_like values(?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert into article table is failed.", err)
		return err
	}
	//sid := strconv.FormatInt(id, 10)
	stmt.QueryRow()
	res, err := stmt.Exec(art.Id, art.UserId, art.ArticleId, int64(1))
	if err != nil {
		sugar.Log.Error("Insert into article_like  is Failed.", err)
		return err
	}
	sugar.Log.Info("Insert into article_like  is successful.")
	l, _ := res.RowsAffected()
	//fmt.Println(" l =", l)
	if l == 0 {
		return errors.New("插入数据失败")
	}
	return nil
	//
}

// 同步取消文章点赞

func SyncArticleShare(db *Sql, value string) error {

	var dl Article
	var art vo.ArticlePlayAddParams
	err := json.Unmarshal([]byte(value), &art)

	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", art)

	//update play num + 1
	stmt, err := db.DB.Prepare("update article set share_num=? where id=?")
	if err != nil {
		sugar.Log.Error("Update  data is failed.The err is ", err)
		return err
	}
	res, err := stmt.Exec(int64(dl.ShareNum+1), art.Id)
	if err != nil {
		sugar.Log.Error("Update  is failed.The err is ", err)
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		sugar.Log.Error("Update  is failed.The err is ", err)
		return err
	}
	if affect == 0 {
		sugar.Log.Error("Update  is failed.The err is ", err)
		return err
	}
	return nil
}

func SyncUserUpdate(db *Sql, value string) error {

	return nil
}

type sycn struct {
}

var Topicmp map[string]*pubsub.Topic

func SyncTopicData(ipfsNode *ipfsCore.IpfsNode, db *Sql, value string) error {
	//监听topic
	topic := "/db-online-sync"
	sugar.Log.Info("开始监听主题 : ", topic)
	sugar.Log.Info("subscrib topic: ", topic)

	ctx := context.Background()
	sugar.Log.Info("加入 主题 房间  : ", topic)
	// 判断 map 是否存在 当前 主题

	tp, err := ipfsNode.PubSub.Join(topic)
	if err != nil {
		sugar.Log.Error("subscribe Join failed.", err)
		return err
	}
	//
	sugar.Log.Info("将tp 加入 到 map中  : ", topic)
	Topicmp = make(map[string]*pubsub.Topic)
	Topicmp["/db-online-sync"] = tp
	sugar.Log.Info("主题map :", Topicmp)

	sugar.Log.Info(" Subscribe topic  tp :", tp)

	sub, err := tp.Subscribe()
	if err != nil {
		sugar.Log.Error("subscribe failed.", err)
		return err
	}

	for {
		data, err := sub.Next(ctx)
		if err != nil {
			sugar.Log.Error("subscribe failed.", err)
			break
		}
		msg := data.Message

		fromId := msg.From
		sugar.Log.Info("来自谁的消息:", string(fromId))
		peerId := ipfsNode.Identity.String()
		sugar.Log.Info("本地节点peerId:", peerId)
		if string(fromId) == peerId {
			sugar.Log.Info("本地节点peerId  等于 本地节点  break ")
			continue
		}

		var recieve vo.SyncMsgParams
		err = json.Unmarshal(msg.Data, &recieve)
		if err != nil {
			sugar.Log.Error("data unmarshal failed", err)
			continue
		}
		sugar.Log.Info("解析收到的消息是:", recieve)

		if recieve.Method == "receiveArticleAdd" {
			//  添加 文章  入库
			sugar.Log.Info("添加 文章  入库:")

		} else if recieve.Method == "receiveArticlePlayAdd" {
			sugar.Log.Info("增加播放次数")

			//  增加播放次数
			//var tmp vo.ChatMsgParams
			//json1, _ := json.Marshal(msg.Data)
			//json.Unmarshal(json1, &tmp)
			//
			//res, err := handleNewMsg(db, tmp)
			//if err != nil {
			//	sugar.Log.Error("handle add message failed.", err)
			//	continue
			//}
			//msg.Data = res
			//jsonStr, _ := json.Marshal(msg)
			//clh.HandlerChat(string(jsonStr))
		} else if recieve.Method == "receiveArticleShareAdd" {
			//  增加 分享 次数
			sugar.Log.Info("增加播放次数:")

		} else if recieve.Method == "receiveUserRegister" {
			// 添加用户 信息
			sugar.Log.Info("添加用户 信息:")

		} else {
			sugar.Log.Info("不满足条件，继续:")
			continue
		}
	}
	return nil
}
