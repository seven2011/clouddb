package mvc

import (
	"encoding/json"
	"errors"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/utils"
	"strconv"
	"time"
)

func AddChatRecord(db *Sql, value string) error {
	var msg ChatRecord
	err := json.Unmarshal([]byte(value), &msg)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return errors.New("解析字段错误")
	}
	sugar.Log.Info("Marshal data is  ", msg)
	id := utils.SnowId()
	t := time.Now().Format("2006-01-02 15:04:05")
	stmt, err := db.DB.Prepare("INSERT INTO chat_record values(?,?,?,?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert into chat_record table is failed.", err)
		return err
	}
	sid := strconv.FormatInt(id, 10)
	stmt.QueryRow()
	res, err := stmt.Exec(sid,msg.RecordName,msg.RecordTalker,msg.RecordImg,msg.CreateBy,t,msg.LastMsg )
	if err != nil {
		sugar.Log.Error("Insert into chat_record  is Failed.", err)
		return err
	}
	sugar.Log.Info("Insert into chat_record  is successful.")
	l, _ := res.RowsAffected()
	if l==0{
		return errors.New("插入记录表数据错误")
	}
	return nil

}