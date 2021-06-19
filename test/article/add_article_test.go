package article

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/cosmopolitann/clouddb/mvc"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/utils"
	"github.com/cosmopolitann/clouddb/vo"
	"strconv"
	"testing"
	"time"
)

func TestAddArticle(t *testing.T) {
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
	//value := `{"id":"444","userId":"409330202166956032","accesstory":"20","accesstoryType":1,"text":"1","tag":"1","playNum":3,"title":"南京","shareNum":4,"thumbnail":"杨洋","fileName":"测试","fileSize":"9"}
//`

//	value:=`{
	//      "accesstory": "QmWQxCg718LJPMfQceXXXTBE9cc67L2M2mQhKxiBmuoQxK,QmY4DpcGLVgwMAWhwbdXjzotV9GYsAePUBUgcjDdYJJfxP,QmY4DpcGLVgwMAWhwbdXjzotV9GYsAePUBUgcjDdYJJfxP,QmcN8AnKciGYHnDz6DNsGG3STUpo71MLqyp4jvxpTVdmBn,QmTVhH34nqxAqhu6QMs8JQi8CGdcWkJcm9xvkx7kGWt6y7,QmbDMypKzWjZQLYo8YvsiBAJ7JK9gWgEH7uL36CQgzr65X,QmTuRRei1RHzGDmtnzs5APtM55LqA3BtGBQobArv4T6i1H",
	//      "accesstoryType": 2,
	//      "text": "记得记得就放假减肥减肥就发烦恼妇女节",
	//      "tag": "关注这些花火大会",
	//      "title": "好多好多话"
	//    }`
	value:=`{
   "id": "325707698052770816",
   "accesstory": "QMabcdefghijk96",
   "text":"正文开始内容",
   "accesstory_type": 0,
   "tag": "标签31",
   "ptime": "1623955566",
  "play_num": 0,
   "share_num": 0,
   "title": "title65",
   "user_id": "323733228975432704",
   "thumbnail": "thumbnail5",
   "file_name": "",
   "file_size": ""
}`
	var art vo.ArticleAddParams
	err = json.Unmarshal([]byte(value), &art)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", art)
	id := utils.SnowId()
	t1:=time.Now().Unix()
	stmt, err := d.Prepare("INSERT INTO article values(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert into article table is failed.", err)
	}
	sid := strconv.FormatInt(id, 10)
	stmt.QueryRow()
	res, err := stmt.Exec(sid, art.UserId, art.Accesstory, art.AccesstoryType, art.Text, art.Tag, t1, 0, 0, art.Title, art.Thumbnail, art.FileName, art.FileSize)
	if err != nil {
		sugar.Log.Error("Insert into article  is Failed.", err)
	}
	l, _ := res.RowsAffected()
	if l == 0 {
	}

	resp := (value)
	fmt.Println("这是返回的数据 =", resp)

}
func Testdb2(sq *sql.DB) mvc.Sql {
	return mvc.Sql{DB: sq}
}
