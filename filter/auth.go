package filter

import (
	"fmt"
	"regexp"

	"github.com/labstack/echo"
)

type (
	//AuthRequestConfig 鉴权配置结构
	AuthRequestConfig struct {
		reg       *regexp.Regexp
		whiteList []string
	}
)

var (
	//DefaultAuthRequestConfig 默认鉴权配置
	DefaultAuthRequestConfig = AuthRequestConfig{
		reg:       regexp.MustCompile(`moonbear\.cn`),
		whiteList: []string{"178.128.115.5"},
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
			if origin == "" && referer == "" {
				return fmt.Errorf("invalide request: both origin and referer are null")
			}
			if origin != "" && !config.reg.MatchString(origin) {
				return fmt.Errorf("invalide request: origin(%v)", origin)
			}
			if referer != "" && config.reg.MatchString(referer) {
				return fmt.Errorf("invalide request: referer(%v)", referer)
			}
			// realIP := c.RealIP()
			// if !utils.StrInSlice(realIP, &config.whiteList) {
			// 	return fmt.Errorf("invalide request: realIP(%v)", realIP)
			// }
			return next(c)
		}
	}
}
