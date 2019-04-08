package danmu

import "testing"

func TestMatchAddress(t *testing.T) {
	bili := &bilibili{}
	ok := bili.MatchAddress("https://live.bilibili.com/1234")
	if !ok {
		t.Error("测试地址匹配失败")
	}
}
