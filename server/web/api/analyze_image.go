package api

import (
	"encoding/json"
	"net/http"
	"net/url"
	"path"

	"github.com/Wisteria30/J-analyzer/lib/recognition"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Hello struct {
	Hello string `json:"Hello"`
}

// 画像の解析用API
func AnalyzeImage() echo.HandlerFunc {
	return func(c echo.Context) error {
		img_url := c.QueryParam("img_url")

		url, err := url.Parse(recognition.RecognitionAPIEndpoint)
		if err != nil {
			logrus.Error("Error cannot create request: ", err)
			return c.JSON(http.StatusBadRequest, err)
		}

		url.Path = path.Join(url.Path, recognition.Recognition)
		q := url.Query()
		q.Set("img_url", img_url)
		url.RawQuery = q.Encode()

		client := new(http.Client)
		req, err := http.NewRequest("GET", url.String(), nil)
		if err != nil {
			logrus.Error("Error cannot create request: ", err)
			return c.JSON(http.StatusBadRequest, err)
		}
		resp, err := client.Do(req)
		if err != nil {
			logrus.Error("Error do not get: ", err)
			return c.JSON(http.StatusBadRequest, err)
		}
		defer resp.Body.Close()

		pai := new(recognition.Paifu)
		if err = json.NewDecoder(resp.Body).Decode(&pai); err != nil {
			logrus.Error("failed to decode a responsed JSON: ", err)
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, pai)
	}
}
