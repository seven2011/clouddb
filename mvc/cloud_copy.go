package mvc

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/utils"
	"github.com/cosmopolitann/clouddb/vo"
	"log"
	"strconv"
	"time"
)

//CopyFile
func CopyFile(db *Sql,value string)(error) {
	//copy file or  copy dir

	var cFile vo.CopyFileParams
	err:=json.Unmarshal([]byte(value), &cFile)
	if err!=nil{
		log.Println("err  ",err)

		return err
	}
	log.Println("解析的数据 file ",cFile)

	//
	//校验 token 是否 满足
	claim,b:=jwt.JwtVeriyToken(cFile.Token)
	if !b{
	}
	sugar.Log.Info("claim := ", claim)
	userid:=claim["UserId"].(string)

	//先查询 数据库 里面 是否 已经存在 相同文件夹
	// 查询
	//step 1  find user info is exist.
	for _,v:=range cFile.Ids{
		//开启事务
		// 查询 id 的 文件 或者 文件夹 信息
		//rows1, _ := db.DB.Query("SELECT * FROM cloud_file where id=?",v)
		//var find File
		//for rows1.Next() {
		//	err := rows1.Scan(&find.Id, &find.UserId, &find.FileName, &find.ParentId, &find.Ptime, &find.FileCid,&find.FileSize,&find.FileType,&find.IsFolder)
		//	if err != nil {
		//		return  err
		//	}
		//}
		log.Println("parent_id = ",cFile.ParentId)
		log.Println("v = ",v)

		rows, err := db.DB.Query("SELECT * from cloud_file as b WHERE (b.file_name,b.user_id,b.is_folder) in (SELECT a.file_name,a.user_id,a.is_folder from cloud_file as a where a.id=?) and b.parent_id=?",v,cFile.ParentId)
		if err!=nil{
			log.Println("errroew = ",err)
			return err
		}
		var s File
		for rows.Next() {
			err := rows.Scan(&s.Id, &s.UserId, &s.FileName, &s.ParentId, &s.Ptime, &s.FileCid,&s.FileSize,&s.FileType,&s.IsFolder)

			if err != nil {
				fmt.Println("query err is ",err)
				return  err
			}
			log.Println("这是查到第一次 多语句查询的 值 file = ",s)

		}
		log.Println("s.id = ",s.Id)
		log.Println("s.id = ",s.Id)

		if s.Id!=""{
			log.Println("文件已经存在")
			return errors.New("文件已经存在")
		}


		log.Println("查到的 文件 数据 结果 = ",s)
		if s.Id==""{
			//0  文件  1 文件夹
			//
			rows, err := db.DB.Query("SELECT * from cloud_file where id=?",v)
			if err!=nil{
				log.Println("errroew = ",err)
				return err
			}
			for rows.Next() {
				err := rows.Scan(&s.Id, &s.UserId, &s.FileName, &s.ParentId, &s.Ptime, &s.FileCid,&s.FileSize,&s.FileType,&s.IsFolder)

				if err != nil {
					fmt.Println("query err is ",err)
					return  err
				}
				log.Println("ss2 = ",s)

			}
				id := utils.SnowId()
				t:=time.Now().Format("2006-01-02 15:04:05")
				stmt, err := db.DB.Prepare("INSERT INTO cloud_file values(?,?,?,?,?,?,?,?,?)")
				if err != nil {
					sugar.Log.Error("Insert into cloud_file table is failed.",err)
					return err
				}
				sid := strconv.FormatInt(id, 10)
				res, err := stmt.Exec(sid, userid, s.FileName, cFile.ParentId,t ,s.FileCid,s.FileSize,s.FileType,1)
				c,_:=res.RowsAffected()
				if c==0{
					return errors.New("插入文件失败")
				}
				fmt.Println("c =",c)
			}
	}

	return nil
}