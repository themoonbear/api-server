package utils

//ErrorOK 错误码
const ErrorOK = 0

const (
	//ErrorParams 错误参数
	ErrorParams = iota + 1000
	//ErrorData 错误数据
	ErrorData
	//ErrorServer 系统错误
	ErrorServer
)

var errorMsg = map[int]string{
	ErrorParams: "params error",
	ErrorData:   "data error",
	ErrorServer: "server error",
}

//GetErrorMsg 获取错误信息
func GetErrorMsg(code int) string {
	if msg, ok := errorMsg[code]; ok {
		return msg
	}
	return Int2String(code)
}

//CatchException 捕获异常
func CatchException(f func(string)) {
	if err := recover(); err != nil {
		switch r := err.(type) {
		case string:
			f(r)
		case error:
			f(r.Error())
		}
	}
}
