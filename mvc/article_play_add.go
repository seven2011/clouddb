package mvc

import (
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	"encoding/json"
	"errors"
)


func ArticlePlayAdd(db *Sql, value string) error {
	var dl Article
	var art vo.ArticlePlayAddParams
	err := json.Unmarshal([]byte(value), &art)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return err
	}
	sugar.Log.Info("Marshal data is  ", art)
	if err != nil {
		sugar.Log.Error("Insert into article table is failed.", err)
		return err
	}
	//select the data is exist.
	rows, err := db.DB.Query("select * from article where id=?", art.Id)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return err
	}
	vl,_:=rows.Columns()
	sugar.Log.Info("vl ", vl)

	for rows.Next() {
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.Title,&dl.ShareNum)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return err
		}

		sugar.Log.Info("Query a entire data is ", dl)
	}
	if dl.Id==""{
      return errors.New(" update is failed .")
	}
	//update play num + 1
	stmt, err := db.DB.Prepare("update article set play_num=? where id=?")
	if err!=nil{
		sugar.Log.Error("Update  data is failed.The err is ", err)
        return err
	}
	res, err := stmt.Exec(int64(dl.PlayNum+1),art.Id)
	if err!=nil{
		sugar.Log.Error("Update  is failed.The err is ", err)
		return err
	}

	affect, err := res.RowsAffected()
	if err!=nil{
		sugar.Log.Error("Update  is failed.The err is ", err)
		return err
	}
	if affect==0{
		sugar.Log.Error("Update  is failed.The err is ", err)
		return err
	}

	return nil
}
// share add

func ArticleShareAdd(db *Sql, value string) error {
	var dl Article
	var art vo.ArticlePlayAddParams
	err := json.Unmarshal([]byte(value), &art)

	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", art)
	if err != nil {
		sugar.Log.Error("Insert into article table is failed.", err)
		return err
	}
	//select the data is exist.
	rows, err := db.DB.Query("select * from article where id=?", art.Id)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return err
	}

	for rows.Next() {
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.ShareNum)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return err
		}

		sugar.Log.Info("Query a entire data is ", dl)
	}
	if dl.Id==""{
		return errors.New(" update is failed .")
	}
	//update play num + 1
	stmt, err := db.DB.Prepare("update article set share_num=? where id=?")
	if err!=nil{
		sugar.Log.Error("Update  data is failed.The err is ", err)
		return err
	}
	res, err := stmt.Exec(int64(dl.ShareNum+1),art.Id)
	if err!=nil{
		sugar.Log.Error("Update  is failed.The err is ", err)
		return err
	}

	affect, err := res.RowsAffected()
	if err!=nil{
		sugar.Log.Error("Update  is failed.The err is ", err)
		return err
	}
	if affect==0{
		sugar.Log.Error("Update  is failed.The err is ", err)
		return err
	}

	return nil
}