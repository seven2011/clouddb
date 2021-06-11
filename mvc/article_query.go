package mvc

import (
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	"encoding/json"
	"errors"
)

//查询文件列表

func ArticleQuery(db *Sql, value string)(data Article,e error) {
	var dl Article

	var list vo.ArticleQueryParams
	err := json.Unmarshal([]byte(value), &list)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return dl,err
	}
	sugar.Log.Info("Marshal data is  ", list)
	// 查询
	rows, err := db.DB.Query("select * from article where id=?", list.Id)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return dl,errors.New("查询下载列表信息失败")
	}
	for rows.Next() {
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType,&dl.Text, &dl.Tag, &dl.Ptime ,&dl.ShareNum,&dl.PlayNum,&dl.Title)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return dl,err
		}
	}
	sugar.Log.Info("Query all data is ", dl)
	return dl,nil

}
