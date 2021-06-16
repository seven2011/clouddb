package mvc

import (
	"encoding/json"
	"errors"
	"github.com/cosmopolitann/clouddb/jwt"
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
	t:=time.Now().Unix()
	//t := time.Now().Format("2006-01-02 15:04:05")
	stmt, err := db.DB.Prepare("INSERT INTO article values(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert into article table is failed.", err)
		return errors.New("插入article 表数据 失败")

	}
	sid := strconv.FormatInt(id, 10)
	stmt.QueryRow()
	res, err := stmt.Exec(sid, art.UserId, art.Accesstory,art.AccesstoryType,art.Text, art.Tag,t , 0,0,art.Title,art.Thumbnail,art.FileName,art.FileSize)
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

func ArticleList(db *Sql, value string) ([]Article,error) {
	var art []Article
	var result vo.ArticleListParams
	err := json.Unmarshal([]byte(value), &result)

	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", result)

	//校验 token 是否 满足
	claim,b:=jwt.JwtVeriyToken(result.Token)
	if !b{
		return art,errors.New("token 失效")
	}
	sugar.Log.Info("claim := ", claim)
	sugar.Log.Error("Marshal data is  result := ", result)
	r:=(result.PageNum-1)*3
	sugar.Log.Info("pageSize := ", result.PageSize)
	sugar.Log.Info("pageNum := ", result.PageNum)
	sugar.Log.Info("r := ", r)
	userid:=claim["UserId"]
	sugar.Log.Info("userid := ", userid)

	//这里 要修改   加上 where  参数 判断

	//todo
	//
	rows, err := db.DB.Query("SELECT * FROM article where user_id =? limit ?,?",userid, r,result.PageSize)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return art, errors.New("查询下载列表信息失败")
	}
	for rows.Next() {
		var dl Article
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.ShareNum,&dl.Title,&dl.Thumbnail,&dl.FileName,&dl.FileSize)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return art, err
		}
		sugar.Log.Info("Query a entire data is ", dl)
		art = append(art, dl)
	}
	if err != nil {
		sugar.Log.Error("Query  article  is Failed.", err)
		return art,err
	}
	sugar.Log.Info("Query  article  is successful.")
	return art,nil

}

