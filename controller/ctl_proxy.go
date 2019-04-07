package controller

import (
	"net/http"

	"api-server/utils"
	"github.com/labstack/echo"
)

type proxyController struct{}

func (ctl proxyController) registerRoute(e *echo.Echo) {
	g := e.Group("/proxy")
	g.GET("/", ctl.proxy)
}

func (ctl proxyController) proxy(ctx echo.Context) error {
	address := ctx.QueryParam("address")
	resData := newResData(ctx)
	if address == "" {
		resData.PackError(utils.ErrorParams)
		return ctx.JSON(http.StatusOK, resData)
	}
	header := &map[string]string{
		"User-Agent":      ctx.Request().Header.Get("User-Agent"),
		"Accept-Language": ctx.Request().Header.Get("Accept-Language"),
		"Accept-Encoding": ctx.Request().Header.Get("Accept-Encoding"),
	}
	body, err := utils.HTTPGet(utils.B64Decode(address), header)
	if err != nil {
		resData.PackError(utils.ErrorServer, err.Error())
		return ctx.JSON(http.StatusOK, resData)
	}
	resData.PackResult(body)
	return ctx.JSON(http.StatusOK, resData)
}
