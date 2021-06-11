package jwt

import (
	"log"
	"testing"
	"time"
)

func TestTtime(t *testing.T){
	value:=time.Now().Unix()
	log.Println("value =",value)


}
