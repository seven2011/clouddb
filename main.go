package main

import (
	"github.com/cosmopolitann/clouddb/mvc"
	"github.com/cosmopolitann/clouddb/sugar"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//日志运行
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	d := mvc.Newdb("")
	e := d.Ping()
	sugar.Log.Info(" 这是 Ping 的 err", e)
	sugar.Log.Info("创建数据库 完成")
	d.DB.Close()
}


