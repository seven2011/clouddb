package mvc

import (
	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	"encoding/json"
	"errors"
)

func DownloadList(db *Sql,value string)(data []DownLoad,e error) {
	//查询下载用户信息
	//userid  就是 sys_user 表的  雪花id
	var d []DownLoad
	var tl vo.TransferListParams
	//解析

	err:=json.Unmarshal([]byte(value), &tl)
	if err!=nil{
		return d ,err
	}
	//

	//校验 token 是否 满足
	claim,b:=jwt.JwtVeriyToken("")
	if !b{
		return d,errors.New("token 失效")
	}
	sugar.Log.Info("claim := ", claim)
	rows, err := db.DB.Query("select * from cloud_transfer where user_id=?", claim["UserId"].(string))
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return d,errors.New("查询下载列表信息失败")
	}
	for rows.Next() {
		var dl DownLoad
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.FileName, &dl.Ptime, &dl.FileCid, &dl.FileSize, &dl.DownPath, &dl.FileType)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return d,err
		}
		sugar.Log.Info("Query a entire data is ", dl)
		d=append(d,dl)
	}
	sugar.Log.Info("Query all data is ", d)
	return d,nil
}
