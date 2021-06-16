package cloud

import (
	"database/sql"
	"fmt"
	"github.com/cosmopolitann/clouddb/sugar"
	"testing"
)

func TestRenameFile(t *testing.T) {
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	//The path is default.
	sugar.Log.Info("Start Open Sqlite3 Database.")
	d, err := sql.Open("sqlite3", "/Users/apple/winter/D-cloud/tables/foo.db")
	if err != nil {
		panic(err)
	}
	sugar.Log.Info("Open Sqlite3 is ok.")
	sugar.Log.Info("Db value is ", d)
	e := d.Ping()
	fmt.Println(" Ping is failed,err:=", e)
	ss := Testdb(d)
	//插入数据
	value := `{
"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI0MDkzMzAyMDIxNjY5NTYwMzIiLCJleHAiOjE2MjU4ODk0NzZ9.OzEFVuB2FcRYurZiii1fpiAqX2KcesfS5arJfVJZQOI",
    "rename":"西游记",
    "id":"411261637169778688",
"isFolder":0,
"parentId":"411193551548846080"
}`
	//b1, e := json.Marshal(fi)
	//fmt.Println(ss)
	//fmt.Println(b1)
	resp := ss.FileRename(string(value))
	fmt.Println("这是返回的数据 =", resp)

}
