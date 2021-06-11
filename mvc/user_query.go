package mvc

import (
	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	"encoding/json"
)

func UserQuery(db *Sql, value string)(data SysUser,e error) {
	var dl SysUser
	var userlist vo.UserListParams
	err := json.Unmarshal([]byte(value), &userlist)

	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", userlist)
	//校验 token 是否 满足
	claim,b:=jwt.JwtVeriyToken(userlist.Token)
	if !b{
		return dl,err
	}
	sugar.Log.Info("claim := ", claim)

	//query
	rows, err := db.DB.Query("select * from sys_user where id=?", userlist.Id)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return dl,err
	}
	for rows.Next() {
		err = rows.Scan(&dl.Id, &dl.PeerId, &dl.Name, &dl.Phone, &dl.Sex, &dl.Ptime, &dl.Utime, &dl.NickName)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return dl,err
		}
		sugar.Log.Info("Query a entire data is ", dl)
	}
	sugar.Log.Info("~~~~~   Delete user into  is Successful ~~~~~~")
	return dl,nil
}

