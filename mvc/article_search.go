package mvc

import (
	"encoding/json"
	"errors"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

//搜索 模糊查询

//todo

func ArticleSearch(db *Sql, value string) (data Article, e error) {
	var dl Article
	var list vo.ArticleQueryParams
	err := json.Unmarshal([]byte(value), &list)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", list)
	// 查询
	rows, err := db.DB.Query("select * from article where id=?", list.Id)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return dl, errors.New("查询下载列表信息失败")
	}
	for rows.Next() {
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.ShareNum)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return dl, err
		}
		sugar.Log.Info("Query a entire data is ", dl)
	}
	sugar.Log.Info("Query all data is ", dl)
	return dl, nil

}
