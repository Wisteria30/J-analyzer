package api

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
)

type Response struct {
	Access_token string
	Token_type   string
	Scope        string
}

func GetAccessToken() echo.HandlerFunc {
	return func(c echo.Context) error {
		code := c.Param("code")
		// TODO セッション由来のstateに変更する
		state := c.Param("state")
		if state != os.Getenv("STATE") {
			logrus.Error("Error do not match state.")
			return c.JSON(http.StatusBadRequest, "Illegal access")
		}

		target_url := fmt.Sprintf(
			"https://api.gyazo.com/oauth/token?client_id=%s&client_secret=%s&redirect_uri=%s&code=%s&grant_type=authorization_code",
			os.Getenv("CLIENT_ID"),
			os.Getenv("CLIENT_SECRET"),
			os.Getenv("CALLBACK_URL"),
			code,
		)
		resp, err := http.Get(target_url)
		if err != nil {
			logrus.Error("Error Request: ", err)
			return c.JSON(http.StatusBadRequest, err)
		}
		// TCPコネクションの終了
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			logrus.Error("Error Response: ", resp.Status)
			return c.JSON(http.StatusBadRequest, resp.Status)
		}

		body, _ := io.ReadAll(resp.Body)
		var res Response
		err = json.Unmarshal(body, &res)
		if err != nil {
			logrus.Error("Error do not read response: ", err)
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, body)
	}
}
