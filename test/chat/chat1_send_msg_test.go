package chat

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"testing"

// 	"github.com/cosmopolitann/clouddb/jwt"
// 	"github.com/cosmopolitann/clouddb/sugar"
// 	"github.com/cosmopolitann/clouddb/vo"

// 	_ "github.com/mattn/go-sqlite3"

// 	shell "github.com/ipfs/go-ipfs-api"
// )

// func TestChatSendMsg(t *testing.T) {
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

// 	req := vo.ChatMsgParams{
// 		RecordId:    "409330202166956089_323733228975432704",
// 		ContentType: 2,
// 		Content:     "content 22222222",
// 		FromId:      "409330202166956089",
// 		ToId:        "323733228975432704",
// 		IsWithdraw:  0,
// 		IsRead:      0,
// 		Ptime:       1623810928,
// 		Token:       token,
// 	}
// 	value, _ := json.Marshal(req)

// 	ss := Testdb(d)
// 	sh := shell.NewShell("localhost:5001")

// 	resp := ss.ChatSendMsg(sh, string(value))
// 	t.Log("获取返回的数据 :=  ", resp)

// }
