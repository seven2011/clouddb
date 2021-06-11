
package cloud

import (
"github.com/cosmopolitann/clouddb/sugar"
"database/sql"
"log"
"testing"
)

func TestTransferAdd(t *testing.T) {
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

//	value:=`{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI0MDkzMzAyMDIxNjY5NTYwMzIiLCJleHAiOjE2MjU4ODk0NzZ9.OzEFVuB2FcRYurZiii1fpiAqX2KcesfS5arJfVJZQOI","parentId":"1"}
//`
	v:=`{
"id":"123",
"fileName":"321",
"fileCid":"Qwer",
"fileSize":123,
"downPath":"23/123",
"fileType":1,
"transferType":1,
"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI0MDkzMzAyMDIxNjY5NTYwMzIiLCJleHAiOjE2MjU4ODk0NzZ9.OzEFVuB2FcRYurZiii1fpiAqX2KcesfS5arJfVJZQOI"
}`
	resp:= ss.TransferAdd(v)
	log.Println("这是返回的数据 =",resp)
}
