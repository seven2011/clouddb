package article

import (
	"database/sql"
	"fmt"
	"github.com/cosmopolitann/clouddb/mvc"
	"github.com/cosmopolitann/clouddb/sugar"
	"testing"
)

func TestAddArticle(t *testing.T) {
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
	ss:= Testdb2(d)

	value:=`{"id":"444","userId":"","accesstory":"20","accesstoryType":1,"text":"1","tag":"1","playNum":3,"title":"成都23","shareNum":4,"thumbnail":"刘亦菲3"}
`
//	value:=`{
//      "accesstory": "QmWQxCg718LJPMfQceXXXTBE9cc67L2M2mQhKxiBmuoQxK,QmY4DpcGLVgwMAWhwbdXjzotV9GYsAePUBUgcjDdYJJfxP,QmY4DpcGLVgwMAWhwbdXjzotV9GYsAePUBUgcjDdYJJfxP,QmcN8AnKciGYHnDz6DNsGG3STUpo71MLqyp4jvxpTVdmBn,QmTVhH34nqxAqhu6QMs8JQi8CGdcWkJcm9xvkx7kGWt6y7,QmbDMypKzWjZQLYo8YvsiBAJ7JK9gWgEH7uL36CQgzr65X,QmTuRRei1RHzGDmtnzs5APtM55LqA3BtGBQobArv4T6i1H",
//      "accesstoryType": 2,
//      "text": "记得记得就放假减肥减肥就发烦恼妇女节",
//      "tag": "关注这些花火大会",
//      "title": "好多好多话"
//    }`

	ss.ArticleAdd(value)

	resp:= (value)
	fmt.Println("这是返回的数据 =",resp)


}
func Testdb2(sq *sql.DB) mvc.Sql {
	return mvc.Sql{DB: sq}
}