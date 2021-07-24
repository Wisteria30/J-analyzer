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
		list, err := gyazo.GetList(*client, nil)
		if err != nil {
			logrus.Error(err)
			return c.JSON(http.StatusBadRequest, err)
		}
		fmt.Println(list.Images)
		return c.JSON(http.StatusOK, *list)
	}
}
