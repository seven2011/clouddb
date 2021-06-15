package article

import (
	"database/sql"
	"github.com/cosmopolitann/clouddb/mvc"
	"github.com/cosmopolitann/clouddb/sugar"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestAddArticleCategory(t *testing.T) {
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	//The path is default.
	sugar.Log.Info("Start Open Sqlite3 Database.")
	d, err := sql.Open("sqlite3", "../../tables/foo.db")
	if err!=nil{
		panic(err)
	}
	sugar.Log.Info("Open Sqlite3 is ok.")
	sugar.Log.Info("Db value is ",d)
	e:=d.Ping()
	sugar.Log.Info(" Ping is failed,err:= ",e)
	ss:=Testdb(d)
	// request json  params
	// test 1
	value:=`{"pageSize":3,"pageNum":0,"accesstoryType":1}
`
	t.Log("request value :=",value)
	resp:= ss.ArticleCategory(value)
	t.Log("result:=",resp)
	// test 2
	//value2:=`{"pageSize":3,"pageNum":1,"accesstoryType":1}
//`
//
//	t.Log("request value :=",value2)
//	resp2:= ss.ArticleCategory(value2)
//	t.Log("result:=",resp2)


}
func Testdb(sq *sql.DB) mvc.Sql {
	return mvc.Sql{DB: sq}
}