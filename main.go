package main

import (
	"github.com/cosmopolitann/clouddb/mvc"
	"github.com/cosmopolitann/clouddb/sugar"
	_ "github.com/mattn/go-sqlite3"
)

type Cloud struct {
	d mvc.Sql
}


func tt() (*Cloud,error){
	//日志运行
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	d := mvc.Newdb("/Users/apple/winter/D-cloud/tables/foo.db")
	e := d.Ping()
	if e!=nil{
		sugar.Log.Info(" 这是 Ping 的 err", e)
		return &Cloud{d:d},e
	}
	sugar.Log.Info("创建数据库 完成")
	return &Cloud{d:d},nil
}



func main(){

	d:=mvc.NTestNode("")
	err:=d.Add()
	sugar.Log.Info("创建数据库 失败 err",err)

}

