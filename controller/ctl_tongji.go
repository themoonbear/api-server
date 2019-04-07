package controller

import (
	"fmt"
	"net/http"

	"api-server/modules/badge"
	"api-server/utils"
	"github.com/labstack/echo"
)

const (
	baiduURL = "https://hm.baidu.com/hm.js?%s"
)

type tjController struct{}

func (ctl tjController) registerRoute(e *echo.Echo) {
	g := e.Group("/tj")
	g.GET("/baidu/:code", ctl.baidu)
}

func (ctl tjController) baidu(ctx echo.Context) error {
	code := ctx.Param("code")
	if code == "" {
		return ctx.String(http.StatusBadRequest, utils.GetErrorMsg(utils.ErrorParams))
	}
	// 通知baidu
	go func() {
		_, err := utils.HTTPGet(fmt.Sprintf(baiduURL, code), nil)
		if err != nil {
			ctx.Logger().Error(err)
		}
	}()
	// 获取徽章
	svg, err := badge.FetchShields()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	ctx.Response().Header().Add("Content-Type", "image/svg+xml; charset=utf-8")
	ctx.Response().Header().Add("Access-Control-Expose-Headers", "Content-Type, Cache-Control, Expires")
	ctx.Response().Header().Add("Cache-Control", "no-cache, no-store, must-revalidate, max-age=0")
	ctx.Response().Header().Add("Expires", "0")
	ctx.Response().Header().Add("Pragma", "no-cache")

	return ctx.String(http.StatusOK, svg)
}
