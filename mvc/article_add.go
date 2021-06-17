package mvc

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/utils"
	"github.com/cosmopolitann/clouddb/vo"
	ipfsCore "github.com/ipfs/go-ipfs/core"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"strconv"
	"time"
)

func AddArticle(ipfsNode *ipfsCore.IpfsNode,db *Sql, value string) error {
	var art vo.ArticleAddParams
	err := json.Unmarshal([]byte(value), &art)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return errors.New("解析字段错误")
	}
	sugar.Log.Info("Marshal data is  ", art)
	id := utils.SnowId()
	t := time.Now().Unix()
	stmt, err := db.DB.Prepare("INSERT INTO article values(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert into article table is failed.", err)
		return errors.New("插入article 表数据 失败")
	}
	sid := strconv.FormatInt(id, 10)
	stmt.QueryRow()
	res, err := stmt.Exec(sid, art.UserId, art.Accesstory, art.AccesstoryType, art.Text, art.Tag, t, 0, 0, art.Title, art.Thumbnail, art.FileName, art.FileSize)
	if err != nil {
		sugar.Log.Error("Insert into article  is Failed.", err)
		return errors.New("插入数据失败")
	}
	l, _ := res.RowsAffected()
	if l == 0 {
		return errors.New("插入数据失败")
	}


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

	err = tp.Publish(ctx,[]byte(value))
	if err != nil {
		sugar.Log.Error("发布错误:", err)
		return err
	}
	sugar.Log.Error("---  发布的消息  完成  ---")

	return nil
}

//

func ArticleList(db *Sql, value string) ([]Article, error) {
	var art []Article
	var result vo.ArticleListParams
	err := json.Unmarshal([]byte(value), &result)

	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", result)

	//校验 token 是否 满足
	claim, b := jwt.JwtVeriyToken(result.Token)
	if !b {
		return art, errors.New("token 失效")
	}
	sugar.Log.Info("claim := ", claim)
	sugar.Log.Error("Marshal data is  result := ", result)
	r := (result.PageNum - 1) * result.PageSize
	sugar.Log.Info("pageSize := ", result.PageSize)
	sugar.Log.Info("pageNum := ", result.PageNum)
	sugar.Log.Info("r := ", r)
	userid := claim["UserId"]
	sugar.Log.Info("userid := ", userid)

	//这里 要修改   加上 where  参数 判断
	//todo
	rows, err := db.DB.Query("SELECT * FROM article where user_id=? limit ?,?",userid,r,result.PageSize)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return art, errors.New("查询下载列表信息失败")
	}
	for rows.Next() {
		var dl Article
		var userId interface{}
		var k=""
		err = rows.Scan(&dl.Id, &userId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.ShareNum, &dl.Title, &dl.Thumbnail, &dl.FileName, &dl.FileSize)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return art, err
		}
		if userId==nil{
			dl.UserId=k
		}else {
			dl.UserId=userId.(string)
		}


		sugar.Log.Info("Query a entire data is ", dl)
		art = append(art, dl)
	}
	if err != nil {
		sugar.Log.Error("Query  article  is Failed.", err)
		return art, err
	}
	sugar.Log.Info("Query  article  is successful.")

	return art, nil

}
