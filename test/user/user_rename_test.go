package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/cosmopolitann/clouddb/mvc"
	"github.com/cosmopolitann/clouddb/sugar"
	"testing"
)

func TestUserRename(t *testing.T) {
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	//The path is default.
	sugar.Log.Info("Start Open Sqlite3 Database.")
	d, err := sql.Open("sqlite3", "/Users/apple/winter/D-cloud/tables/foo.db")
	if err!=nil{
		panic(err)
	}
	sugar.Log.Info("Open Sqlite3 is ok.")
	sugar.Log.Info("Db value is ",d)
	e:=d.Ping()
	ss:= Testdb(d)
	//插入数据
	var fi = mvc.File{
		Id:         "1",
		UserId:     "408217533556985856",
		FileName:   "红楼梦",
		ParentId:   "0",
		FileCid:    "Qmcid",
		FileSize:   100,
		FileType:   11,
		IsFolder:   0,
		Ptime:      1232,
	}
	b1, e := json.Marshal(fi)
	fmt.Println(e)
	fmt.Println(b1)

	//这里 改成 穿 json 字符串，字段 要改成更新之后的数据。

	//{"id":"4324","peerId":"124","name":"20","phone":1,"sex":"1","nickName":"nick"}
	value:=`{"phone":"1888"}`
	//resp:= ss.UserAdd(string(b1)

	resp:= ss.UserLogin(value)

	fmt.Println("这是返回的数据 =",resp)


}