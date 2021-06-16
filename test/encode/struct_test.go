package encode

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cosmopolitann/clouddb/mvc"
	"github.com/cosmopolitann/clouddb/utils"
	"strconv"
	"testing"
	"time"
)

type SQl struct {
	Db *sql.DB
	All
}

type All struct {
	Chat
	Article123
}

type Chat struct {
	ADb *sql.DB
}

type Article123 struct {
	ADb *sql.DB
}

func (c *Chat) ChatMsg() {

}
func (c *Chat) ChatSend() {
	//查询数据
	fmt.Println("--------")
	fmt.Println("--------mvc =", c.ADb)

	e := AddArticleTest(c.ADb, `{"Id":"4324","UserId":"124","Accesstory":"20","AccesstoryType":1,"Text":"1","Tag":"1","PlayNum":3,"ShareNum":4}`)
	if e != nil {

	}

}
func (c *Article123) ArticleAdd() {
}

func NewTestStrcut() *All {
	return &All{
		Chat:       Chat{ADb: InitDB1("")},
		Article123: Article123{ADb: InitDB1("")},
	}
}

//func Newdb(path string)Sql{
//	return Sql{DB: InitDB(path)}
//}
func InitDB1(path string) *sql.DB {
	//
	//mvc, err := sql.Open("sqlite3", path)
	db, err := sql.Open("sqlite3", "/Users/apple/winter/D-cloud/tables/foo.mvc")
	checkErr(err)
	fmt.Println("Db value is ", db)

	return db
}
func checkErr(err error) {
	if err != nil {
		fmt.Println("The connection to the database failed.")
		panic(err)
	}
}

func TestStruct(t *testing.T) {
	//
	s := NewTestStrcut()
	fmt.Println("mvc = ", s.Chat.ADb)

	s.Chat.ChatSend()

}
func AddArticleTest(d *sql.DB, value string) error {

	var art mvc.Article
	err := json.Unmarshal([]byte(value), &art)

	if err != nil {
		fmt.Println(" err =", err)

	}

	id := utils.SnowId()
	t := time.Now().Format("2006-01-02 15:04:05")

	stmt, err := d.Prepare("INSERT INTO article values(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
		fmt.Println(" err =", err)

	}
	sid := strconv.FormatInt(id, 10)
	stmt.QueryRow()
	res, err := stmt.Exec(sid, art.UserId, art.Accesstory, art.AccesstoryType, art.Text, art.Tag, t, art.PlayNum, art.ShareNum)
	if err != nil {
		return err
	}

	l, _ := res.RowsAffected()
	fmt.Println(" l =", l)
	if l == 0 {
		return errors.New("插入数据失败")
	}

	return nil

}
