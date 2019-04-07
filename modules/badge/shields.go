package badge

import (
	"api-server/utils"
)

const (
	shieldsKey = "Shields:baidu"
	shieldsURL = "https://img.shields.io/badge/%E7%99%BE%E5%BA%A6-%E7%BB%9F%E8%AE%A1-red.svg?logo=baidu"
)

// FetchShields 获取徽章
func FetchShields() (string, error) {
	data, ok := utils.GetCache(shieldsKey)
	if ok {
		return data, nil
	}
	data, err := utils.HTTPGet(shieldsURL, nil)
	if err != nil {
		return "", err
	}
	utils.AddCache(shieldsKey, data)
	return data, nil
}
