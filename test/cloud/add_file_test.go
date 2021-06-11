package cloud

import (
	"github.com/cosmopolitann/clouddb/mvc"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestAddFile(t *testing.T) {
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
	var fi = vo.CloudAddFileParams{
		Id:         "409054047794892800",
		FileName:   "三国演义",
		ParentId:   "1",
		FileCid:    "Qmcid12312",
		FileSize:   100,
		FileType:   0,
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiIxODg4IiwiZXhwIjoxNjI1ODc4MzYxfQ.FkABVsrygfxKq2_GWP5pG2G9oYpUqD1yw5cA3boB-Dc",
	}

	b1, e := json.Marshal(fi)
	fmt.Println(e)
	fmt.Println("这是 json 数据",string(b1))



	resp:= ss.AddFile(string(b1))
	fmt.Println("这是返回的数据 =",resp)


}
func Testdb(sq *sql.DB) mvc.Sql {
	return mvc.Sql{DB: sq}
}
