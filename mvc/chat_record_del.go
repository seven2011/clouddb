package mvc

import (
	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	"encoding/json"
	"errors"
)

func ChatRecordDel(db *Sql, value string) error {

	var rdel vo.ChatRecordDelParams
	err := json.Unmarshal([]byte(value), &rdel)
	if err != nil {
		return err
	}

	claim,b:=jwt.JwtVeriyToken(rdel.Token)
	if !b{
		return errors.New("token 失效")
	}
	sugar.Log.Info("claim := ", claim)
	stmt, err := db.DB.Prepare("delete from chat_record where id=?")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(rdel.Id)
	if err != nil {
		sugar.Log.Error("delete  chat_record data  is failed.", err)
		return err
	}
	c, _ := res.RowsAffected()
	if c == 0 {
		return err
	}
	sugar.Log.Info("delete record is successful.")
	return nil

}
