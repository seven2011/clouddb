package mvc

import (
	"encoding/json"
	"errors"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/utils"
	"github.com/cosmopolitann/clouddb/vo"
	"strconv"
	"time"
)

func AddArticle(db *Sql, value string) error {
	var art vo.ArticleAddParams
	err := json.Unmarshal([]byte(value), &art)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return errors.New("解析字段错误")
	}
	sugar.Log.Info("Marshal data is  ", art)
	id := utils.SnowId()
	t := time.Now().Format("2006-01-02 15:04:05")
	stmt, err := db.DB.Prepare("INSERT INTO article values(?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert into article table is failed.", err)
		return errors.New("插入article 表数据 失败")

	}
	sid := strconv.FormatInt(id, 10)
	stmt.QueryRow()
	res, err := stmt.Exec(sid, art.UserId, art.Accesstory,art.AccesstoryType, art.Text, art.Tag,t , 0,0,art.Title,art.Thumbnail)
	if err != nil {
		sugar.Log.Error("Insert into article  is Failed.", err)
		return errors.New("插入数据失败")
	}
	sugar.Log.Info("Insert into article  is successful.")
	l, _ := res.RowsAffected()
	if l==0{
		return errors.New("插入数据失败")
	}
	return nil
}
