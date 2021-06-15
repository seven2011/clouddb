package cloud

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	"testing"
)

func TestAddFolder(t *testing.T) {
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
		FileName: "月亮湾",
		ParentId: "411193551548846080",
		Token:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI0MTExOTI3MjA0NDc1MDg0ODAiLCJleHAiOjE2MjYzMzM1NzJ9.EDv7k__JEugp-57RU3wDsJ1Swa3C0t-Ofr4KeQQhzeA",
	}

	b1, e := json.Marshal(f1)
	fmt.Println(e)
	fmt.Println(string(b1))
	resp:= ss.AddFolder(string(b1))
	fmt.Println("这是返回的数据 =",resp)


}