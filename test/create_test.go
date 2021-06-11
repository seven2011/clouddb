package main

import (
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestCreateFile(t *testing.T){
	f, err := os.OpenFile("../token.txt", os.O_CREATE|os.O_RDWR,0600)
	defer f.Close()
	if err !=nil {
		fmt.Println(err.Error())
	} else {
		//_,err=f.Write([]byte("123"))
		//if err!=nil{
		//	t.Log("err ",err)
		//}
		bytes, err := ioutil.ReadFile("../token.txt")
		if err != nil {
			log.Fatal(err)
		}
		t.Log("内容===",string(bytes))

	}
	//

	err = os.Remove("../token.txt")
	if err!=nil{
		t.Log(err)
	}

}

func TestResp(t *testing.T){
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	err:=errors.New("这是错误")
	l:=vo.ResponseErrorMsg(400,err.Error())
	t.Log("l =",l)


}