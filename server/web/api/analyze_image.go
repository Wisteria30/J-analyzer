package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"

	"github.com/Wisteria30/J-analyzer/lib/analyzer"
	"github.com/Wisteria30/J-analyzer/lib/recognition"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

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
		pai.SyantenType = 1
		pai.Flag = 63
		pai.MeldedBlocks = []int{}

		pai_json, err := json.Marshal(pai)
		if err != nil {
			logrus.Error("Error before post analyzer: ", err)
			return c.JSON(http.StatusBadRequest, err)
		}
		req, err = http.NewRequest("POST", analyzer.AnalyzerAPIEndpoint, bytes.NewBuffer(pai_json))
		if err != nil {
			logrus.Error("Error cannot create request: ", err)
			return c.JSON(http.StatusBadRequest, err)
		}
		resp, err = client.Do(req)
		if err != nil {
			logrus.Error("Error do not post: ", err)
			return c.JSON(http.StatusBadRequest, err)
		}
		defer resp.Body.Close()
		fmt.Println("RESP JSON")
		fmt.Println(resp.Body)

		result := new(analyzer.Result)
		if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
			logrus.Error("failed to decode a responsed JSON: ", err)
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, result)
	}
}
