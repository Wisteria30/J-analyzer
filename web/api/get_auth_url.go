package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	// "net/url"
	"os"
)

type URL struct {
	Redirect_url string `json:"redirect_url"`
}

func GetAuthURL() echo.HandlerFunc {
	return func(c echo.Context) error {

		redirect_url := fmt.Sprintf(
			"https://api.gyazo.com/oauth/authorize?client_id=%s&redirect_uri=%s&response_type=code&state=%s",
			os.Getenv("CLIENT_ID"),
			os.Getenv("CALLBACK_URL"),
			os.Getenv("STATE"),
		)
		u := URL{redirect_url}
		return c.JSON(http.StatusOK, u)
	}
}
