package mvc

import (
	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	"encoding/json"
	"errors"
	"log"
)

//递归删除文件夹下面所有的文件或者文件夹
var  delArray []string
func del(db *Sql, parent_id string,userId string) {
	//如果有子文件夹，则：
	if parent_id != "" {
		log.Println("parent_id = ",parent_id)
		var f File
		var f1 []File
		rows, _ := db.DB.Query("SELECT * FROM cloud_file where parent_id=? and user_id=?", parent_id,userId)
		for rows.Next() {
			err := rows.Scan(&f.Id, &f.UserId, &f.FileName, &f.ParentId, &f.Ptime, &f.FileCid, &f.FileSize, &f.FileType, &f.IsFolder)
			if err != nil {
				log.Println("find err is ", err)
			}
			if f.Id != "" {
					f1 = append(f1, f)
			}
		}
		for i:=0;i<len(f1);i++{
			delArray = append(delArray, f1[i].Id)
			if f1[i].IsFolder == 1 {
				del(db, f1[i].Id,userId)
			}
		}
	}
	log.Println("打印 最终 的 数组 结果 delArray =-= ", delArray)
}
func Delete(db *Sql, value string) error {
	//test
	// parent_id ==  408293247874502656
	//var parent_id string="408293247874502656"
	//var delMap = []string{}
	//del(db, parent_id)
	var d vo.CloudDeleteParams
	err := json.Unmarshal([]byte(value), &d)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("解析数据  是 ",d)
	//验证token 是否满足条件
	//校验 token 是否 满足
	claim,b:=jwt.JwtVeriyToken(d.Token)
	if !b{
		return err
	}
	sugar.Log.Info("claim := ", claim)
	//如果 是  文件夹  就 递归删除  如果 是文件 就直接删除
	//查询 id 判断类型
		for _,v:=range d.Ids{
		rows, err := db.DB.Query("select * from cloud_file where id=?",v)
		if err != nil {
			sugar.Log.Error("Query data is failed.Err is ", err)
			return errors.New("查询下载列表信息失败")
		}
		var dl File
		for rows.Next() {
			err = rows.Scan(&dl.Id, &dl.UserId, &dl.FileName, &dl.ParentId,&dl.Ptime, &dl.FileCid, &dl.FileSize, &dl.FileType,&dl.IsFolder)
			if err != nil {
				sugar.Log.Error("Query scan data is failed.The err is ", err)
				return err
			}
		}
		if dl.IsFolder==1{
			del(db, dl.ParentId,claim["UserId"].(string))
		}
		delArray=append(delArray,string(v))
		}

		//删除 所有的 id
		// 开启事务
	log.Println(" =========== 数组 信息 ",delArray)

	tx,err:=db.DB.Begin()
		if err!=nil{
			return errors.New("删除错误")
		}

		for _,v:=range delArray{
			stmt, err := db.DB.Prepare("delete from cloud_file where id=?")
			if err != nil {
				sugar.Log.Info("删除文件失败，错误err:= ",err)
				tx.Rollback()
				return errors.New("删除错误")
			}

			res, err := stmt.Exec(v)
			if err != nil {
				sugar.Log.Error("删除文件失败，错误err:= ",err)
				tx.Rollback()

				return errors.New("删除错误")
			}
			log.Println(res)
		}
		err=tx.Commit()
		if err!=nil{
			sugar.Log.Info("删除文件失败，错误err:= ",err)
			return errors.New("删除错误")
		}
	    log.Println("数组 ：=",delArray)
	return nil

}
