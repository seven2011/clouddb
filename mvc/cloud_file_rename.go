package mvc

import (
	"encoding/json"
	"errors"
	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

func CloudFileRename(db *Sql, value string) error {

	var art vo.CloudFileRenameParams
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
	//userid:=claim["UserId"].(string)
	sugar.Log.Info("claim := ", claim)
	//查询数据
	stmt, err := db.DB.Prepare("UPDATE cloud_file set file_name=? where id=?")
	if err!=nil{
		sugar.Log.Error("update cloud_file is failed.Err is ", err)
		return errors.New("更新数据失败")
	}
	res, err := stmt.Exec(art.Rename,art.Id)
	if err!=nil{
		sugar.Log.Error("update cloud_file is failed.Err is ", err)
		return err
	}
	affect, err := res.RowsAffected()
	if affect==0{
		sugar.Log.Error("update cloud_file is failed.Err is ", err)
		return err
	}
	return nil

}
