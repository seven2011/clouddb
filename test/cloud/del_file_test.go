package cloud

import (
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	"database/sql"
	"encoding/json"
	"fmt"
	"testing"
)

func TestAddDelete(t *testing.T) {
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
	var f1 =vo.CloudAddFolderParams{
		Id:       "123",
		FileName: "文学巨著",
		ParentId: "12",
		Token:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiIxODg4IiwiZXhwIjoxNjI1ODc4MzYxfQ.FkABVsrygfxKq2_GWP5pG2G9oYpUqD1yw5cA3boB-Dc",
	}
	b1, e := json.Marshal(f1)
	fmt.Println(e)
	fmt.Println(string(b1))

	value:=`{"ids":["409292472099803136","409302447882768384"]}
`


	resp:= ss.DeleteAll(value)
	fmt.Println("这是返回的数据 =",resp)
}