package user

import (
	"database/sql"
	"fmt"
	"github.com/cosmopolitann/clouddb/sugar"
	"testing"
)

func TestUserRename(t *testing.T) {
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
	ss := Testdb(d)

/*
	Id       string `json:"id"`
	Rename   string `json:"rename"`
	IsFolder int64  `json:"isFolder"`
	Token    string `json:"token"`
	ParentId string `json:"parentId"`
 */
	value := `{"id":"411826072401743872","rename":"月亮与六便士","isFolder":"0","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI0MDkzMzAyMDIxNjY5NTYwMzIiLCJleHAiOjE2MjU4ODk0NzZ9.OzEFVuB2FcRYurZiii1fpiAqX2KcesfS5arJfVJZQOI","parentId":"123"}`

	resp := ss.FileRename(value)

	fmt.Println("这是返回的数据 =", resp)

}
