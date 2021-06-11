package mvc

import (
	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/utils"
	"github.com/cosmopolitann/clouddb/vo"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)


func ChatRecordAdd(db *Sql, value string)(string, error) {

	var record vo.ChatRecoredAddParams
	err := json.Unmarshal([]byte(value), &record)

	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return "",err
	}
	sugar.Log.Info("Marshal data is  ", record)
	//
	//校验 token 是否 满足
	claim,b:=jwt.JwtVeriyToken(record.Token)
	if !b{
		return "", errors.New("token 失效")

	}
	sugar.Log.Info("claim := ", claim)

	id := utils.SnowId()
	t := time.Now().Format("2006-01-02 15:04:05")
	stmt, err := db.DB.Prepare("INSERT INTO chat_record values(?,?,?,?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert into chat_msg table is failed.", err)
		return "",err
	}
	sid := strconv.FormatInt(id, 10)
	stmt.QueryRow()
	res, err := stmt.Exec(sid,record.RecordName,record.RecordImg, record.RecordTalker,record.CreateBy,t,record.LastMsg)
	if err != nil {
		sugar.Log.Error("Insert into chat_msg  is Failed.", err)
		return "",err
	}
	sugar.Log.Info("Insert into chat_msg  is successful.")
	l, _ := res.RowsAffected()
	fmt.Println(" l =", l)
	if l==0{
		return "",errors.New("插入chat_msg数据失败")
	}
	return sid ,nil

}
