package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GyazoからImages一覧を取得する
func GetHealth() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "status OK.")
	}
}
