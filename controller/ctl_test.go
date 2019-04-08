package controller

import (
	"testing"
)

func TestInit(t *testing.T) {
	e := Init()
	if e == nil {
		t.Error("测试初始化失败")
	}
}
