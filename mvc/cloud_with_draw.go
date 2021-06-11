package mvc

import (
	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	"encoding/json"
	"errors"
)

// 撤回消息


func ChatWithDraw(db *Sql, value string) error {
	var art vo.ChatMsgWithDrawParams
	err := json.Unmarshal([]byte(value), &art)

	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", art)
	//查询是否存在记录

	//校验 token 是否 满足
	claim,b:=jwt.JwtVeriyToken(art.Token)
	if !b{
		return errors.New("token 失效")
	}
	sugar.Log.Info("claim := ", claim)

	//查询数据
	stmt, err := db.DB.Prepare("UPDATE chat_msg set is_with_draw=? where id=?")
	if err!=nil{
		sugar.Log.Error("update article_like is failed.Err is ", err)
		return err
	}

	//  0 未撤回   1  撤回
	res, err := stmt.Exec(int64(1),art.Id)
	if err!=nil{
		sugar.Log.Error("update article_like is failed.Err is ", err)
		return err
	}
	affect, err := res.RowsAffected()
	if affect==0{
		sugar.Log.Error("update article_like is failed.Err is ", err)
		return err
	}

	return nil

}