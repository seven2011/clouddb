package vo

import (
	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
)

func tt(){

	//校验 token 是否 满足
	claim,b:=jwt.JwtVeriyToken("")
	if !b{
		//return errors.New("token 失效")
	}
	sugar.Log.Info("claim := ", claim)


}