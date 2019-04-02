package controller

import (
	"net/http"

	"github.com/labstack/echo"
	pm "github.com/themoonbear/api-server/modules/danmu"
	"github.com/themoonbear/api-server/utils"
)

type danmuController struct{}

func (ctl danmuController) registerRoute(e *echo.Echo) {
	g := e.Group("/danmu")
	g.GET("/:address", ctl.danmu)
}

func (ctl danmuController) danmu(ctx echo.Context) error {
	address := ctx.Param("address")
	resData := newResData(ctx)
	if address == "" {
		resData.PackError(utils.ErrorParams)
		return ctx.JSON(http.StatusOK, resData)
	}
	data, err := pm.ParseAddress(utils.B64Decode(address))
	if err != nil {
		resData.PackError(utils.ErrorServer, err.Error())
		return ctx.JSON(http.StatusOK, resData)
	}
	resData.PackResult(data)
	return ctx.JSON(http.StatusOK, resData)
}
