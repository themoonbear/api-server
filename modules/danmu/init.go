package danmu

import (
	"errors"
)

type danmu interface {
	MatchAddress(address string) bool
	ParseAddress(address string) (data interface{}, err error)
}

// ParseAddress 解析地址
func ParseAddress(address string) (data interface{}, err error) {
	for _, v := range store {
		if v.MatchAddress(address) {
			return v.ParseAddress(address)
		}
	}
	return nil, errors.New("不支持该平台")
}

var store []danmu

func registerParser(pareser danmu) {
	store = append(store, pareser)
}
