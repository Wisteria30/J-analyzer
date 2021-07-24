package api

import (
	"net/http"

	"github.com/Wisteria30/J-analyzer/lib/gyazo"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// Gyazoからユーザを取得する
func GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("access_token")
		if err != nil {
			logrus.Error("Cookie Error: ", err)
			return c.JSON(http.StatusBadRequest, err)
		}
		client, err := gyazo.NewClient(cookie.Value)
		if err != nil {
			logrus.Error("Error Request: ", err)
			return c.JSON(http.StatusBadRequest, err)
		}
		user, err := gyazo.GetUser(*client)
		if err != nil {
			logrus.Error(err)
			return c.JSON(http.StatusBadRequest, err)
		}
		logrus.Info(user)
		return c.JSON(http.StatusOK, *user)
	}
}
