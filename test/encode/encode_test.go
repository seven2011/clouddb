package encode

import (
	"encoding/json"
	"fmt"
	"testing"
)

type TestEncoding struct {
	Id   int
	Name string
}
type TestEncoding1 struct {
	D []TestEncoding
}

func TestEncode(t *testing.T) {
	//
	value := `
{
"d":[{"id":1,"name":"nick"}]
}`
	var en TestEncoding1
	// 将字符串反解析为结构体
	json.Unmarshal([]byte(value), &en)
	fmt.Println(" 解析 之后的值 ：= ", en)
}
