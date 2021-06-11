package utils

import "testing"

func TestSnowId(t *testing.T) {
	id := SnowId()
	t.Log("生成的雪花id=", id)
}
