package chat

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"testing"

// 	"github.com/cosmopolitann/clouddb/sugar"
// 	"github.com/cosmopolitann/clouddb/vo"

// 	_ "github.com/mattn/go-sqlite3"

// 	"github.com/cosmopolitann/clouddb/jwt"
// 	shell "github.com/ipfs/go-ipfs-api"
// )

// func TestChatCreateRecord(t *testing.T) {
// 	sugar.InitLogger()
// 	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
// 	//The path is default.
// 	sugar.Log.Info("Start Open Sqlite3 Database.")
// 	d, err := sql.Open("sqlite3", "/data/projects/clouddb/tables/foo.db")
// 	if err != nil {
// 		panic(err)
// 	}
// 	sugar.Log.Info("Open Sqlite3 is ok.")
// 	sugar.Log.Info("Db value is ", d)
// 	err = d.Ping()
// 	if err != nil {
// 		panic(err)
// 	}

// 	token, _ := jwt.GenerateToken("409330202166956089", 30*24*60*60)

// 	fmt.Println(token)

// 	req := vo.ChatRecordParams{
// 		Name:    "Record Name 2222",
// 		Img:     "https://xxxx/img.png",
// 		FromId:  "409330202166956089",
// 		ToId:    "323733228975432704",
// 		Ptime:   1623825227,
// 		LastMsg: "nothing 22222",
// 		Token:   token,
// 	}

// 	value, _ := json.Marshal(req)

// 	ss := Testdb(d)
// 	sh := shell.NewShell("localhost:5001")
// 	resp := ss.ChatCreateRecord(sh, string(value))
// 	t.Log("获取返回的数据 :=  ", resp)

// }
