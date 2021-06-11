package mvc

import (
	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

//MoveFile
func MoveFile(db *Sql,value string)(error) {
	//move file or  copy dir
	var mvFile vo.MoveFileParams
	err:=json.Unmarshal([]byte(value), &mvFile)
	if err!=nil{
		sugar.Log.Info("Unmarshal 参数值 ： ", err)

		return err
	}
	//step 1  find user info is exist.

	//校验 token 是否 满足
	claim,b:=jwt.JwtVeriyToken(mvFile.Token)
	if !b{
		return errors.New("token 失效")
	}
	sugar.Log.Info("解析token 参数值 ： ", claim)

	userid:=claim["UserId"].(string)


	for _,v:=range mvFile.Ids{
		//开启事务
		rows1, _ := db.DB.Query("SELECT * FROM cloud_file where id=?",v)
		var m File

		for rows1.Next() {
			err := rows1.Scan(&m.Id, &m.UserId, &m.FileName, &m.ParentId, &m.Ptime, &m.FileCid,&m.FileSize,&m.FileType,&m.IsFolder)

			if err != nil {
				fmt.Println("query err is ",err)
				return  err
			}
		}
		rows, _ := db.DB.Query("SELECT * FROM cloud_file where file_name=? and parent_id=? and user_id=?",m.FileName,mvFile.ParentId,userid)
		var s File

		for rows.Next() {
			err := rows.Scan(&s.Id, &s.UserId, &s.FileName, &s.ParentId, &s.Ptime, &s.FileCid,&s.FileSize,&s.FileType,&s.IsFolder)

			if err != nil {
				fmt.Println("query err is ",err)
				return  err
			}
		}

		log.Println(" 这是查找的结果 s =",s)
		if s.Id!=""{
			log.Println("文件夹已经存在")
			return errors.New("文件夹已经存在")
		}
		if s.Id==""{

			//0  文件  1 文件夹
		//	if s.IsFolder==0{

				//stmt, err := mvc.DB.Prepare("INSERT INTO cloud_file values(?,?,?,?,?,?,?,?,?,?)")
				stmt, err := db.DB.Prepare("UPDATE cloud_file set parent_id=? where id=?")

				//update userinfo set username=? where uid=?
				fmt.Println(" 这是 需要 更新的 id == ",m.Id)
				fmt.Println(" 这是 需要 更新的 Pid == ",mvFile.ParentId)

				res, err := stmt.Exec(mvFile.ParentId, m.Id)
				if err != nil {
					sugar.Log.Error("Update cloud_file table is failed.",err)
					return err
				}
				c,_:=res.RowsAffected()
				//if c==0{
				//
				//}
				fmt.Println("更新  c =",c)

			}
			//if s.IsFolder==1{
			//	// 更新 文件
			//
			//	//stmt, err := mvc.DB.Prepare("INSERT INTO cloud_file values(?,?,?,?,?,?,?,?,?,?)")
			//	stmt, err := db.DB.Prepare("UPDATE cloud_file set parent_id=? where id=?")
			//
			//	//update userinfo set username=? where uid=?
			//	res, err := stmt.Exec(mvFile.ParentId, s.Id)
			//	if err != nil {
			//		sugar.Log.Error("Update cloud_file table is failed.",err)
			//		return err
			//	}
			//
			//	c,_:=res.RowsAffected()
			//	//if c==0{
			//	//
			//	//}
			//	fmt.Println("更新 记录 c =",c)
			//}
		}



	return nil
}