package filter

import (
	"testing"
)

func TestCheckWhileList(t *testing.T) {
	data := "localhost"
	ok := checkWhileList(data, &DefaultAuthRequestConfig.whiteList)
	if ok {
		t.Error("测试验证白名单错误")
	}
}
