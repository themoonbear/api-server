package controller

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/themoonbear/utils"
	"net/http"
)

func Init() *echo.Echo {
	e := echo.New()
	initMiddleWare(e)
	initRouter(e)
	return e
}

func initMiddleWare(echo *echo.Echo) {
	echo.Pre(middleware.CORS())
	echo.Use(middleware.RequestID())
	echo.Use(middleware.Recover())
}

func initRouter(echo *echo.Echo) {
	echo.GET("/proxy/:address", proxy)
}

func proxy(ctx echo.Context) error {
	address := ctx.Param("address")
	resData := utils.NewResData(ctx)
	if address == "" {
		resData.PackError(utils.ErrorParams)
		return ctx.JSON(http.StatusOK, resData)
	}
	body, err := utils.HttpGet(ctx, utils.B64Decode(address))
	if err != nil {
		resData.PackError(utils.ErrorServer, err.Error())
		return ctx.JSON(http.StatusOK, resData)
	}
	resData.PackResult(body)
	return ctx.JSON(http.StatusOK, resData)
}
