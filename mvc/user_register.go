package mvc

import (
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/utils"
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

func AddUser(db *Sql,value string)(error){
	//user string ==> user struct
	//Add sys_user
	//create snow id

	var user SysUser
	err:=json.Unmarshal([]byte(value), &user)
	if err!=nil{

	}
	sugar.Log.Info("params ：= ",user)

	l,e:= FindIsExistUser(db,user)
	if e!=nil{
		sugar.Log.Error("FindIsExistUser info is Failed.")
	}
	// l > 0 user is exist.
	sugar.Log.Error("-----------1")

	if l>0{
		sugar.Log.Error("user is exist.")
		return errors.New("user is exist.")
	}
	sugar.Log.Error("-----------2")


	//inExist insert into sys_user.
	id := utils.SnowId()
	//create now time
	t:=time.Now().Format("2006-01-02 15:04:05")
	stmt, err := db.DB.Prepare("INSERT INTO sys_user values(?,?,?,?,?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert data to sys_user is failed.")
		return  err
	}
	sid := strconv.FormatInt(id, 10)
	res, err := stmt.Exec(sid, user.PeerId, user.Name, user.Phone, user.Sex, t, t,user.NickName)
	if err != nil {
		sugar.Log.Error("Insert data to sys_user is failed.",res)
		return err
	}
	c,_:=res.RowsAffected()
	sugar.Log.Info("~~~~~   Insert into sys_user data is Successful ~~~~~~",c)
	//生成 token
	// 手机号
	//token,err:=jwt.GenerateToken(user.Phone,60)

	return nil
}

func FindIsExistUser(db *Sql,user SysUser)(int64,error){
	var s SysUser
	sugar.Log.Info("start sys_user is exist local user info.")
	sugar.Log.Info("user info is ",user.Phone)
	sugar.Log.Info("user info is ",user)

	rows, _ := db.DB.Query("SELECT * FROM sys_user where phone=?",user.Phone)
	for rows.Next() {
		err := rows.Scan(&s.Id, &s.PeerId, &s.Name, &s.Phone, &s.Sex, &s.Ptime, &s.Utime, &s.NickName)
		if err!=nil{
			sugar.Log.Error(" query is failed. ",err)

			return 0, err
		}
		sugar.Log.Info(" user info is ",s)
	}
	//is exist
	sugar.Log.Info(" FindOne data is ",s.Id)

	if s.Id!=""{
	return 1,nil
	}else{
		return 0,nil
	}

}