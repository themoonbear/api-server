package controller

import (
	"api-server/utils"
	"github.com/labstack/echo"
)

//ResData 返回数据结构体
type ResData struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
	Error  string      `json:"error"`
	Trace  string      `json:"trace"`
}

func newResData(ctx echo.Context) *ResData {
	data := &ResData{
		Trace: getRequestTrace(ctx),
	}
	ctx.Set("res_data", data)
	return data
}

func getRequestTrace(ctx echo.Context) string {
	if ctx == nil {
		return ""
	}
	return ctx.Response().Header().Get(echo.HeaderXRequestID)
}

//PackError 打包错误数据
func (res *ResData) PackError(code int, msg ...interface{}) {
	res.Code = code
	if len(msg) == 0 {
		res.Error = utils.GetErrorMsg(code)
	} else {
		res.Error = utils.Struct2String(msg)
	}
}

//PackResult 打包结果数据
func (res *ResData) PackResult(result interface{}) {
	res.Code = utils.ErrorOK
	res.Result = result
}
