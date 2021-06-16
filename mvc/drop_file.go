package mvc

import (
	"github.com/cosmopolitann/clouddb/sugar"
	"encoding/json"
	"fmt"
)

func DeleteOneFile(db *Sql,value string)error {

	var dFile DeleteParams
	err:=json.Unmarshal([]byte(value), &dFile)
	if err!=nil{

	}
	fmt.Println(" 这是 解析后的 数据 dFile=== id",dFile.DropFile[0].Id)
	// 多个 回滚
	fmt.Println(" 这是 解析后的 数据 dFile=== ")

	for _,v:=range dFile.DropFile{
		tx,_:=db.DB.Begin()

		stmt, err := db.DB.Prepare("delete from sys_user where id=?")
		checkErr(err)
		fmt.Println(" 这是 解析后的 数据 v.id=== ",v.Id)
		res, err := stmt.Exec(v.Id)
		if err != nil {
			sugar.Log.Error("Insert into cloud_file table is failed.",err)
			//rowback
			tx.Rollback()
			return err
		}
		c,_:=res.RowsAffected()
		if c==0{
			tx.Rollback()
		}
		fmt.Println(res)
		tx.Commit()
	}
	sugar.Log.Info("Insert into file  is successful.")
	return nil
}