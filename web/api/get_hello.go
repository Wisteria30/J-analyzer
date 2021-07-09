package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GyazoからImages一覧を取得する
func GetHello() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Access TokenをCookieに保存する
		cookie := new(http.Cookie)
		cookie.Name = "hello"
		cookie.Value = "hello world"
		c.SetCookie(cookie)
		return c.String(http.StatusOK, "Hello World")
	}
}
