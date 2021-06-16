package mvc

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/utils"
	"github.com/cosmopolitann/clouddb/vo"
	"strconv"
	"time"
)


func ChatRecordAdd(db *Sql, value string)(ChatRecord, error) {
	var resp ChatRecord

	var record vo.ChatRecoredAddParams
	err := json.Unmarshal([]byte(value), &record)

	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return resp,err
	}
	sugar.Log.Info("Marshal data is  ", record)
	//
	//校验 token 是否 满足
	claim,b:=jwt.JwtVeriyToken(record.Token)
	if !b{
		return resp, errors.New("token 失效")

	}
	sugar.Log.Info("claim := ", claim)
	var sid string
	//0 用传进来的  1  生成新的
	if record.IsActive==0{
		sid=record.Id
	}
	if record.IsActive==1{
		id := utils.SnowId()
		sid= strconv.FormatInt(id, 10)
	}
	sugar.Log.Info("  雪花 Id  = ", sid)


	t := time.Now().Format("2006-01-02 15:04:05")
	stmt, err := db.DB.Prepare("INSERT INTO chat_record values(?,?,?,?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert into chat_msg table is failed.", err)
		return resp,err
	}
	stmt.QueryRow()
	res, err := stmt.Exec(sid,record.Name,record.Img, record.FromId,t,record.LastMsg,record.ToId)
	if err != nil {
		sugar.Log.Error("Insert into chat_msg  is Failed.", err)
		return resp,err
	}
	sugar.Log.Info("Insert into chat_msg  is successful.")
	l, _ := res.RowsAffected()
	fmt.Println(" l =", l)
	if l==0{
		return resp,errors.New("插入chat_msg数据失败")
	}

	//
	resp.Id=sid
	resp.Name=record.Name
	resp.Img=record.Img
	resp.FromId=record.FromId
	resp.Toid=record.ToId
	resp.Ptime=time.Now()
	resp.LastMsg=record.LastMsg
	// json

	//b1, e := json.Marshal(resp)
	//if e!=nil{
	//	return "",errors.New("解析失败")
	//}

	return resp ,nil
}
