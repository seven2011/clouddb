package cloud

import (
	"github.com/cosmopolitann/clouddb/sugar"
	"database/sql"
	"fmt"
	"testing"
)

//DropFile
func TestDropFile(t *testing.T) {
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	//The path is default.
	sugar.Log.Info("Start Open Sqlite3 Database.")
	d, err := sql.Open("sqlite3", "/Users/apple/winter/D-cloud/tables/foo.mvc")
	if err!=nil{
		panic(err)
	}
	sugar.Log.Info("Open Sqlite3 is ok.")
	sugar.Log.Info("Db value is ",d)
	e:=d.Ping()
	fmt.Println(" Ping is failed,err:=",e)
	ss:= Testdb(d)
	//插入数据
	value:=`{
    "DropFile":[
        {
            "Id":"408283148502175744",
           "UserId":"4082175335569858561",
"FileName":"红楼梦",
 "ParentId":"1",
 "Ptime":"12",
            "FileCid":"Qmcid",
            "FileSize":20,
"FileStatus":1,
            "FileType":1,
"IsFolder":0
        }
    ]
}`
	//b1, e := json.Marshal(fi)
	//fmt.Println(ss)
	//fmt.Println(b1)
	resp:= ss.DeleteOneFile(string(value))
	fmt.Println("这是返回的数据 =",resp)


}