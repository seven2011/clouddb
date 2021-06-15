package article

import (
	"database/sql"
	"github.com/cosmopolitann/clouddb/sugar"
	"testing"
)

func TestArticleAboutMe(t *testing.T) {
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
	value:=`{"pageSize":3,"pageNum":0,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI0MDkzMzAyMDIxNjY5NTYwMzIiLCJleHAiOjE2MjU4ODk0NzZ9.OzEFVuB2FcRYurZiii1fpiAqX2KcesfS5arJfVJZQOI"}
`
	t.Log("request value :=",value)
	resp:= ss.ArticleAboutMe(value)
	t.Log("result:=",resp)


}
