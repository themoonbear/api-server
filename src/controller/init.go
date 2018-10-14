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

func initMiddleWare(e *echo.Echo) {
	e.Pre(middleware.CORS())
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
}

func initRouter(e *echo.Echo) {
	e.GET("/proxy", proxy)
}

func proxy(ctx echo.Context) error {
	address := ctx.QueryParam("address")
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
