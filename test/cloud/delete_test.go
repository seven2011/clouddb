package cloud

import (
	"database/sql"
	"fmt"
	"github.com/cosmopolitann/clouddb/sugar"
	"testing"
)

func TestDelete(t *testing.T) {
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
//
	value:=`{"Id":"4324","UserId":"124","Accesstory":"20","AccesstoryType":1,"Text":"1","Tag":"1","PlayNum":3,"ShareNum":4}
`
//	resp:= ss.DeleteOneFile(value)
//
//	fmt.Println("这是返回的数据 =",resp)

	resp1:=ss.DeleteAll(value)

	fmt.Println("这是返回的数据 =",resp1)

}