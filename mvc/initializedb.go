package mvc

import (
	"database/sql"
	"github.com/cosmopolitann/clouddb/sugar"
)

type Sql struct {
	DB *sql.DB
}

type NewTestNode struct {
	db Sql
}

func NTestNode(path string)(*NewTestNode){
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	sugar.Log.Info("创建数据库 完成")
	//
	sql:=Newdb(path)

	return &NewTestNode{db: sql}
}

func (n *NewTestNode)Add()error{
	//
	err:=n.db.Ping()
	if err!=nil{
		sugar.Log.Error("打开 数据库 失败 err is ",err)
	}
	return err
}
//
func (n *NewTestNode)UserRegister(value string)string{
	//
	data:=n.db.UserRegister(value)
	return data
}

func (n *NewTestNode)UserLogin(value string)string{
	//
	data:=n.db.UserLogin(value)
	return data
}



func Newdb(path string) Sql {
	sugar.Log.Info("---- ===== ------")

	return Sql{DB: InitDB(path)}
}

func InitDB(path string)(*sql.DB){
	//
	//mvc, err := sql.Open("sqlite3", path)
	if path==""{
		path="../tables/foo.db"
	}
	sugar.Log.Info(" 数据库路径  = ",path)
	sugar.Log.Info("Start Open Sqlite3 Database.")
	db, err := sql.Open("sqlite3", path)
	checkErr(err)
	sugar.Log.Info("Open Sqlite3 is ok.")
	sugar.Log.Info("Db value is ",db)

	return db
}
func checkErr(err error) {
	if err != nil {
		sugar.Log.Error("The connection to the database failed.")
		panic(err)
	}
}