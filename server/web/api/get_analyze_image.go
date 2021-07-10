package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// 画像の解析用API
func GetAnalyzeImage() echo.HandlerFunc {
	return func(c echo.Context) error {
		url := c.Param("url")
		if url == "" {
			logrus.Error("Error do not analyze.")
			return c.JSON(http.StatusBadRequest, "unmatch requests")
		}
		return c.String(http.StatusOK, url)
	}
}
