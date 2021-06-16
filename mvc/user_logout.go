package mvc

import (
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	"encoding/json"
)

func UserDel(db *Sql, value string) error {
	var userdel vo.UserDelParams
	err := json.Unmarshal([]byte(value), &userdel)
	if err != nil {
		sugar.Log.Error("解析参数失败:", err)
	}
	sugar.Log.Info("解析参数信息是:", userdel)
	//delete
	stmt, err := db.DB.Prepare("delete from sys_user where id=?")
	if 	err!=nil{
		sugar.Log.Error("Delete user info is failed.",err)
		return err
	}
	//token
	//claim,b:=jwt.JwtVeriyToken(userdel.Token)
	//if !b{
	//	return err
	//}
	//sugar.Log.Info("claim := ", claim)

	sugar.Log.Info("userId is  ",userdel.Id)
	res, err := stmt.Exec(userdel.Id)
	if err != nil {
		sugar.Log.Error("Delete user info is  failed.",err)
		return err
	}
	c,_:=res.RowsAffected()
	if c==0{
		sugar.Log.Error("Delete user info is  failed.",err)
		return err
	}
	sugar.Log.Info("~~~~~   Delete user into  is Successful ~~~~~~", c)
	return nil
}

