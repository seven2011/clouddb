package mvc

import (
	"encoding/json"
	"errors"
	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

func UserLogin(db *Sql,value string) (string,error) {
   	//解析传进来的参数信息

	var userLogin vo.UserLoginParams
	err := json.Unmarshal([]byte(value), &userLogin)

	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", userLogin)
	//先查询数据库是否已经注册   如果 未注册 请注册  注册了 生成token
	//查询数据库
	c,err,user:=FindIsExistLoginUser(db,userLogin.Phone)
	if err!=nil{
		return "",err
	}
	if c==0{
		return "",errors.New("请先注册用户")
	}

	////生成 token
	//	// 手机号
	token,err:=jwt.GenerateToken(user.Id,30*24*60*60)

	////  验证 jwt
	//
	//_, flag,b := jt.GetClaim(userLogin.Token)
	//if flag != jt.TOKEN_ERR_LEN && flag != jt.TOKEN_ERR_EXPIRED {
	//
	//}
	//if b{
	//	return errors.New(" Token 过期")
	//}
	//查询数据库  根据 用户手机号   查出用户信息 返回

	//struct => json
	var resp vo.UserLoginRespParams

	resp.Token=token
	resp.UserInfo=user
	sugar.Log.Info(" 获取的 返回信息= ", resp)

	strUser, err := json.Marshal(resp)
	     if err != nil {
			 sugar.Log.Error("Marshal is failed.Err is ", err)
			return "",errors.New("解析参数失败")
		 }
	return string(strUser) ,nil
}
func FindIsExistLoginUser(db *Sql,data string)( int64,error,vo.RespSysUser){
	var s vo.RespSysUser
	sugar.Log.Info("start sys_user is exist local user info.")
	sugar.Log.Info("user info is ",data)

	rows, _ := db.DB.Query("SELECT * FROM sys_user where phone=?",data)
	for rows.Next() {
		err := rows.Scan(&s.Id, &s.PeerId, &s.Name, &s.Phone, &s.Sex, &s.Ptime, &s.Utime, &s.NickName)
		if err!=nil{
			sugar.Log.Error(" query is failed. ",err)

			return 0, err,s
		}
		sugar.Log.Info(" user info is ",s)
	}
	//is exist
	sugar.Log.Info(" FindOne data is ",s.Id)

	if s.Id!=""{
		//说明 有 这个用户

		return 1,nil,s
	}else{
		//说明 没有 这个用户
		return 0,nil,s
	}

}
