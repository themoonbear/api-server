package utils

import (
	"fmt"
	"strconv"

	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

//Struct2String 结构转字符串
func Struct2String(src interface{}) string {
	b, err := json.Marshal(src)
	if err != nil {
		panic(err)
	}
	return string(b)
}

//Int2String int 转字符串
func Int2String(v int) string {
	return strconv.Itoa(v)
}

//Int642String int64 转字符串
func Int642String(v int64) string {
	return strconv.FormatInt(v, 10)
}

//String2Int 字符串转 int
func String2Int(str string) int {
	i, e := strconv.Atoi(str)
	if e != nil {
		panic(fmt.Errorf("%s convert to int error :%s", str, e.Error()))
	}
	return i
}

//StrInSlice 集合包含目标字符串
func StrInSlice(dst string, arry *[]string) bool {
	for _, item := range *arry {
		if item == dst {
			return true
		}
	}
	return false
}
