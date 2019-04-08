package badge

import "testing"

func TestFetchShields(t *testing.T) {
	svg, err := FetchShields()
	if svg == "" || err != nil {
		t.Error("测试获取徽章失败")
	}
}
