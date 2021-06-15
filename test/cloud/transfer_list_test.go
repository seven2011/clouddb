package cloud

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	"testing"
)

func TestCloudTransferList(t *testing.T) {
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
	fmt.Println(" Ping is failed,err:=",e)
	ss:= Testdb(d)
	//插入数据

	//var f1 =vo.CloudAddFolderParams{
	//	Id:       "123",
	//	FileName: "文学巨著",
	//	ParentId: "12",
	//	Token:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiIxODg4IiwiZXhwIjoxNjI1ODc4MzYxfQ.FkABVsrygfxKq2_GWP5pG2G9oYpUqD1yw5cA3boB-Dc",
	//}

	//test
	var tt =vo.TransferListParams{Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI0MDkzMzAyMDIxNjY5NTYwMzIiLCJleHAiOjE2MjU4ODk0NzZ9.OzEFVuB2FcRYurZiii1fpiAqX2KcesfS5arJfVJZQOI"}

	b1, e := json.Marshal(tt)
	fmt.Println(e)
	fmt.Println(string(b1))
	resp:= ss.TransferList(string(b1))
	fmt.Println("这是返回的数据 =",resp)


}