package controller

import (
	"api-server/filter"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Init 初始化控制器
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
	e.Use(filter.AuthRequest())
}

func initRouter(e *echo.Echo) {
	new(proxyController).registerRoute(e)
	new(danmuController).registerRoute(e)
	new(tjController).registerRoute(e)
}
