package mvc

import (
	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	"encoding/json"
	"errors"
)

//查询文件列表

func TransferList(db *Sql, value string)(data []DownLoad,e error) {
	var list vo.TransferListParams
	err := json.Unmarshal([]byte(value), &list)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", list)
	// 查询


	//校验 token 是否 满足
	claim,b:=jwt.JwtVeriyToken(list.Token)
	if !b{
	}
	sugar.Log.Info("claim := ", claim)

	var arrfile []DownLoad
	rows, err := db.DB.Query("select * from cloud_transfer where user_id=?", claim["UserId"].(string))
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return arrfile,errors.New("查询下载列表信息失败")
	}
	for rows.Next() {
		var dl DownLoad
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.FileName,&dl.Ptime, &dl.FileCid, &dl.FileSize, &dl.DownPath,&dl.FileType,&dl.TransferType)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return arrfile,err
		}
		sugar.Log.Info("Query a entire data is ", dl)
		arrfile=append(arrfile,dl)
	}
	sugar.Log.Info("Query all data is ", arrfile)
	return arrfile,nil

	sugar.Log.Info("Insert into article  is successful.")

	return arrfile,nil

}
