package filter

import (
	"fmt"
	"regexp"

	"github.com/labstack/echo"
)

type (
	//AuthRequestConfig 鉴权配置结构
	AuthRequestConfig struct {
		whiteList []*regexp.Regexp
	}
)

var (
	//DefaultAuthRequestConfig 默认鉴权配置
	DefaultAuthRequestConfig = AuthRequestConfig{
		whiteList: []*regexp.Regexp{
			regexp.MustCompile(`moonbear\.cn`),
			regexp.MustCompile(`iwlist\.github\.io`),
		},
	}
)

//AuthRequest 鉴权请求
func AuthRequest() echo.MiddlewareFunc {
	return AuthRequestWithConfig(DefaultAuthRequestConfig)
}

//AuthRequestWithConfig 鉴权请求使用配置
func AuthRequestWithConfig(config AuthRequestConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			origin := c.Request().Header.Get("Origin")
			referer := c.Request().Header.Get("Referer")
			// if origin == "" && referer == "" {
			// 	return fmt.Errorf("invalide request: both origin and referer are null")
			// }
			if origin != "" && !checkWhileList(origin, &config.whiteList) {
				return fmt.Errorf("invalide request: origin(%v)", origin)
			}
			if referer != "" && !checkWhileList(referer, &config.whiteList) {
				return fmt.Errorf("invalide request: referer(%v)", referer)
			}
			return next(c)
		}
	}
}

func checkWhileList(dst string, whiteList *[]*regexp.Regexp) bool {
	for _, reg := range *whiteList {
		if reg.MatchString(dst) {
			return true
		}
	}
	return false
}
