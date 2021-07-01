package api

import (
	"fmt"
	"net/http"

	"github.com/Wisteria30/J-analyzer/lib/gyazo"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// GyazoからImages一覧を取得する
func GetImages() echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("token")
		if err != nil {
			logrus.Error("Error Request: ", err)
			return c.JSON(http.StatusBadRequest, err)
		}
		client, err := gyazo.NewClient(cookie.Value)
		if err != nil {
			logrus.Error("Error Request: ", err)
			return c.JSON(http.StatusBadRequest, err)
		}
		fmt.Println(client)
		return c.String(http.StatusOK, cookie.Value)
	}
}
