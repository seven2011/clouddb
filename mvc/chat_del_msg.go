package mvc

import (
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	"encoding/json"
)

//删除消息

func ChatMsgDel(db *Sql,value string)error {
	var msgdel vo.ChatMsgDelParams
	err:=json.Unmarshal([]byte(value), &msgdel)
	if err!=nil{
		return err
	}
		stmt, err := db.DB.Prepare("delete from chat_msg where id=?")
		if err!=nil{
			return err
		}
		res, err := stmt.Exec(msgdel.Id)
		if err != nil {
			sugar.Log.Error("delete data from chat_msg  table is failed.",err)
			return err
		}
		c,_:=res.RowsAffected()
		if c==0{
			sugar.Log.Error("delete data from chat_msg  table is failed.",err)
			return err
		}
	sugar.Log.Info("delete data from chat_msg  table is failed")
	return nil

}