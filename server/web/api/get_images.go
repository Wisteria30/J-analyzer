package api

import (
	"fmt"
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/Wisteria30/J-analyzer/lib/gyazo"
	"github.com/Wisteria30/J-analyzer/middlewares"
	"github.com/Wisteria30/J-analyzer/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// GyazoからImages一覧を取得する
func GetImages() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		token := c.Get("auth").(*auth.Token)

		user := models.User{}
		dbs.DB.Table("users").Where(models.User{UID: token.UID}).First(&user)
		client, err := gyazo.NewClient(user.AccessToken)
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
