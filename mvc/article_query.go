package mvc

import (
	"encoding/json"
	"errors"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

//查询文件列表

func ArticleQuery(db *Sql, value string)(data vo.ArticleResp,e error) {
	//var dl Article
	var dl vo.ArticleResp
	var list vo.ArticleQueryParams
	err := json.Unmarshal([]byte(value), &list)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return dl,err
	}
	sugar.Log.Info("Marshal data is  ", list)
	// 查询
	//rows, err := db.DB.Query("select * from article where id=?", list.Id)


	rows, err := db.DB.Query("SELECT a.*,b.peer_id,b.name,b.phone,b.sex,b.nickname from article as a LEFT JOIN sys_user as b on a.user_id=b.id where a.user_id=(select c.user_id from article as c where id =?) and a.id=?;", list.Id,list.Id)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return dl,errors.New("查询下载列表信息失败")
	}
	for rows.Next() {
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType,&dl.Text, &dl.Tag, &dl.Ptime ,&dl.ShareNum,&dl.PlayNum,&dl.Title,&dl.Thumbnail,&dl.PeerId,&dl.Name,&dl.Phone,&dl.Sex,&dl.NickName)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return dl,err
		}
	}

	sugar.Log.Info("Query all data is ", dl)
	return dl,nil

}
