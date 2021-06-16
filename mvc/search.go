package mvc

import (
	"encoding/json"
	"errors"
	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	"strconv"
)

//查询文件列表

func Search(db *Sql, value string) (data []File, e error) {
	var s vo.SearchFileParams
	var arrfile []File

	err := json.Unmarshal([]byte(value), &s)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", s)
	// 查询

	//校验 token 是否 满足
	claim, b := jwt.JwtVeriyToken(s.Token)
	if !b {
		return arrfile, errors.New("token 失效")
	}
	var or string
	if s.Order == "" {
		or = "ptime"
	}
	if s.Order == "time" {
		or = "ptime"
	}
	if s.Order == "name" {
		or = "file_name"

	}
	if s.Order == "type" {
		or = "file_type"

	}
	if s.Order == "size" {
		or = "file_size"

	}
	sugar.Log.Info("排序方式:", or)

	userid := claim["UserId"].(string)
	sugar.Log.Info("claim := ", claim)
	sugar.Log.Info("UserId := ", userid)

	sql := "select * from cloud_file where user_id= ? and file_name like'%" + s.Content + "%'" + " order by " + or
	rows, err := db.DB.Query(sql, userid)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return arrfile, errors.New("查询下载列表信息失败")
	}
	for rows.Next() {
		var dl File
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.FileName, &dl.ParentId, &dl.Ptime, &dl.FileCid, &dl.FileSize, &dl.FileType, &dl.IsFolder)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return arrfile, err
		}
		sugar.Log.Info("Query a entire data is ", dl)
		arrfile = append(arrfile, dl)
	}
	sugar.Log.Info("Query all data is ", arrfile)
	return arrfile, nil

	sugar.Log.Info("Search  cloud_file  is successful.")

	return arrfile, nil

}

// 文章查询

func ARticleSearch(db *Sql, value string) (data []Article, e error) {
	var s vo.ArticleSearchParams
	var arrfile []Article

	err := json.Unmarshal([]byte(value), &s)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", s)
	// 查询

	////校验 token 是否 满足
	//claim,b:=jwt.JwtVeriyToken(s.Token)
	//if !b{
	//	return arrfile,errors.New("token 失效")
	//}

	//userid:=claim["UserId"].(string)
	//sugar.Log.Info("claim := ", claim)
	//sugar.Log.Info("UserId := ", userid)
	r := (s.PageNum - 1) * 3

	//

	str := strconv.FormatInt(r, 10)
	pageSize := strconv.FormatInt(s.PageSize, 10)

	sql := "select * from article where title like'%" + s.Title + "%' limit " + str + "," + pageSize
	rows, err := db.DB.Query(sql)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return arrfile, errors.New("查询下载列表信息失败")
	}
	for rows.Next() {
		var dl Article
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.ShareNum, &dl.Title, &dl.Thumbnail, &dl.FileName, &dl.FileSize)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return arrfile, err
		}
		sugar.Log.Info("Query a entire data is ", dl)
		arrfile = append(arrfile, dl)
	}
	sugar.Log.Info("Query all data is ", arrfile)
	return arrfile, nil

	sugar.Log.Info("Insert into article  is successful.")

	return arrfile, nil

}
