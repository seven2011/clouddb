package mvc

import (
	"encoding/json"
	"errors"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	"time"
)

func SyncUser(db *Sql,value string)(error){
	
	var user SysUser
	err:=json.Unmarshal([]byte(value), &user)
	if err!=nil{
	sugar.Log.Error("解析失败:",err)
		return err
	}
	sugar.Log.Info("params ：= ",user)
	//l,e:= FindIsExistUser(db,user)
	//if e!=nil{
	//	sugar.Log.Error("FindIsExistUser info is Failed.")
	//}
	//// l > 0 user is exist.
	//sugar.Log.Error("-----------1")
	//
	//if l>0{
	//	sugar.Log.Error("user is exist.")
	//	return errors.New("user is exist.")
	//}

	//inExist insert into sys_user.
//	id := utils.SnowId()
	//create now time
	//t:=time.Now().Format("2006-01-02 15:04:05")
	t:=time.Now().Unix()
	stmt, err := db.DB.Prepare("INSERT INTO sys_user values(?,?,?,?,?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert data to sys_user is failed.")
		return  err
	}

	//sid := strconv.FormatInt(user.Id, 10)
	res, err := stmt.Exec(user.Id, user.PeerId, user.Name, user.Phone, user.Sex, t, t,user.NickName)
	if err != nil {
		sugar.Log.Error("Insert data to sys_user is failed.",res)
		return err
	}
	c,_:=res.RowsAffected()
	sugar.Log.Info("~~~~~   Sync into sys_user data is Successful ~~~~~~",c)
	//生成 token
	// 手机号
	//token,err:=jwt.GenerateToken(user.Phone,60)

	return nil
}

// 文章

func SyncArticle(db *Sql,value string)(error) {
	var art vo.ArticleAddParams
	err := json.Unmarshal([]byte(value), &art)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return errors.New("解析字段错误")
	}
	sugar.Log.Info("Marshal data is  ", art)
	//id := utils.SnowId()
	//t := time.Now().Format("2006-01-02 15:04:05")
	t:=time.Now().Unix()

	stmt, err := db.DB.Prepare("INSERT INTO article values(?,?,?,?,?,?,?,?,?,?)")

	if err != nil {
		sugar.Log.Error("Insert into article table is failed.", err)
		return errors.New("插入article 表数据 失败")

	}
	//sid := strconv.FormatInt(id, 10)
	stmt.QueryRow()
	res, err := stmt.Exec(art.Id, art.UserId, art.Accesstory,art.AccesstoryType, art.Text, art.Tag,t , 0, art.Title,0)
	if err != nil {
		sugar.Log.Error("Insert into article  is Failed.", err)
		return errors.New("插入数据失败")
	}
	sugar.Log.Info("Insert into article  is successful.")
	l, _ := res.RowsAffected()
	if l==0{
		return errors.New("插入数据失败")
	}
	return nil

}

//

func SyncAticlePlay(db *Sql,value string)(error) {
	//更新字段  is_ike = 1
	var art vo.SyncArticleGiveLikeParams
	err := json.Unmarshal([]byte(value), &art)

	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", art)
	//插入新的一条记录
	//id := utils.SnowId()
	stmt, err := db.DB.Prepare("INSERT INTO article_like values(?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert into article table is failed.", err)
		return err
	}
	//sid := strconv.FormatInt(id, 10)
	stmt.QueryRow()
	res, err := stmt.Exec(art.Id, art.UserId,art.ArticleId,int64(1))
	if err != nil {
		sugar.Log.Error("Insert into article_like  is Failed.", err)
		return err
	}
	sugar.Log.Info("Insert into article_like  is successful.")
	l, _ := res.RowsAffected()
	//fmt.Println(" l =", l)
	if l==0{
		return errors.New("插入数据失败")
	}
	return nil
	//
}

// 同步取消文章点赞

func SyncArticleShare(db *Sql,value string)(error) {

		var dl Article
		var art vo.ArticlePlayAddParams
		err := json.Unmarshal([]byte(value), &art)

		if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
		sugar.Log.Info("Marshal data is  ", art)

		//update play num + 1
		stmt, err := db.DB.Prepare("update article set share_num=? where id=?")
		if err!=nil{
		sugar.Log.Error("Update  data is failed.The err is ", err)
		return err
	}
		res, err := stmt.Exec(int64(dl.ShareNum+1),art.Id)
		if err!=nil{
		sugar.Log.Error("Update  is failed.The err is ", err)
		return err
	}

		affect, err := res.RowsAffected()
		if err!=nil{
		sugar.Log.Error("Update  is failed.The err is ", err)
		return err
	}
		if affect==0{
		sugar.Log.Error("Update  is failed.The err is ", err)
		return err
	}
		return nil
}

func SyncUserUpdate(db *Sql,value string)(error) {
     return nil
}
