package mvc

import (
	"encoding/json"
	"fmt"
	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

func TransferDel(db *Sql, value string) error {

	var dFile vo.TransferDelParams
	err := json.Unmarshal([]byte(value), &dFile)
	if err != nil {
		return err
	}
	//
	//校验 token 是否 满足
	claim, b := jwt.JwtVeriyToken(dFile.Token)
	if !b {
		return err
	}
	sugar.Log.Info("claim := ", claim)

	for _, v := range dFile.Ids {
		tx, _ := db.DB.Begin()
		//  这里 表名 要改一下
		//todo
		stmt, err := db.DB.Prepare("delete from cloud_transfer where id=?")
		if err != nil {
			return err
		}
		res, err := stmt.Exec(v)
		if err != nil {
			sugar.Log.Error("Insert into cloud_file table is failed.", err)
			//rowback
			tx.Rollback()
			return err
		}
		c, _ := res.RowsAffected()
		if c == 0 {
			tx.Rollback()
		}
		fmt.Println(res)
		tx.Commit()
	}
	sugar.Log.Info("Insert into file  is successful.")

	return nil

}
