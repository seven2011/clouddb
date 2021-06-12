package cloud

import (
	"database/sql"
	"github.com/cosmopolitann/clouddb/sugar"
	"log"
	"testing"
)


func TestArticleSearch(t *testing.T) {
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
	value:=`{"pageSize":3,"pageNum":1,"title":"成"}
`
	resp:= ss.ArticleSearch(value)
	log.Println("这是返回的数据 =",resp)
}
