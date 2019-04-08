package utils

import "testing"

func TestAddCache(t *testing.T) {
	data := "123"
	key := "test:key"
	AddCache(key, data)
	exp, ok := GetCache(key)
	if exp != data || !ok {
		t.Error("测试添加缓存失败")
	}
}
