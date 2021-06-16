package chat

// import (
// 	"database/sql"
// 	"fmt"
// 	"testing"
// 	"time"

// 	"github.com/cosmopolitann/clouddb/sugar"

// 	"github.com/cosmopolitann/clouddb/jwt"
// 	_ "github.com/mattn/go-sqlite3"

// 	shell "github.com/ipfs/go-ipfs-api"
// )

// func TestChatListenMsg(t *testing.T) {
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

// 	ss := Testdb(d)

// 	sh := shell.NewShell("localhost:5001")

// 	token, _ := jwt.GenerateToken("409330202166956089", 30*24*60*60)

// 	fmt.Println(token)

// 	var cl ChatLister

// 	resp := ss.ChatListenMsg(sh, token, &cl)
// 	t.Log("获取返回的数据 :=  ", resp)

// 	time.Sleep(time.Hour)

// }

// type ChatLister struct{}

// func (cl *ChatLister) HandlerChat(abc string) {
// 	fmt.Println("1111", abc, "2222")
// }
