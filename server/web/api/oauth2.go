package api

import (
	"context"
	"net/http"
	"time"

	"github.com/Wisteria30/J-analyzer/lib/gyazo"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type TokenRequest struct {
	Code  string `query:"code"`
	State string `query:"state"`
}

func GetAuthCode() echo.HandlerFunc {
	return func(c echo.Context) error {
		config := gyazo.GetConnect()
		url := config.AuthCodeURL("hoge")
		return c.Redirect(http.StatusFound, url)
	}
}

func GetToken() echo.HandlerFunc {
	return func(c echo.Context) error {
		t := new(TokenRequest)
		if err := c.Bind(t); err != nil {
			logrus.Error("Error Request: ", err)
			return c.JSON(http.StatusBadRequest, err)
		}
		config := gyazo.GetConnect()
		context := context.Background()

		token, err := config.Exchange(context, t.Code)
		if err != nil {
			logrus.Error("Error Request: ", err)
			return c.JSON(http.StatusBadRequest, err)
		}
		if !token.Valid() {
			logrus.Error("Invalid token: ", err)
			return c.JSON(http.StatusBadRequest, err)
		}

		// Access TokenをCookieに保存する
		// cookie := new(http.Cookie)
		// cookie.Name = "token"
		// cookie.Value = token.AccessToken
		// cookie.Expires = time.Now().Add(24 * time.Hour)
		// cookie.HttpOnly = true
		// cookie.Secure = true
		c.SetCookie(&http.Cookie{
			Name:     "access_token",
			Value:    token.AccessToken,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
			Domain:   "http://localhost:8170/",
			SameSite: 4,
			Secure:   true,
		})
		// リダイレクト先にCookieを保存する
		// c.Response().After(func() {
		// 	c.SetCookie(cookie)
		// })
		return c.Redirect(http.StatusFound, "http://localhost:8170/")
	}
}
