package cloud

import (
	"github.com/cosmopolitann/clouddb/sugar"
	"database/sql"
	"log"
	"testing"
)

func TestFileList(t *testing.T) {
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
	log.Println(" Ping is failed,err:=",e)
	ss:=Testdb(d)
	value:=`{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiIxODg4IiwiZXhwIjoxNjI1ODc4MzYxfQ.FkABVsrygfxKq2_GWP5pG2G9oYpUqD1yw5cA3boB-Dc","parentId":"1"}
`
	resp:= ss.FileList(value)
	log.Println("这是返回的数据 =",resp)
}
