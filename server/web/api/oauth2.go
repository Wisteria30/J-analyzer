package api

import (
	"context"
	"crypto/rand"
	"errors"
	"net/http"

	"github.com/Wisteria30/J-analyzer/lib/gyazo"
	"github.com/Wisteria30/J-analyzer/middlewares"
	"github.com/Wisteria30/J-analyzer/models"
	"gorm.io/gorm"

	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type AuthResponse struct {
	Url string `json:"url"`
	State string `json:"state"`
}

type TokenRequest struct {
	Code  string `query:"code"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func stateGenerator(digit uint32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// 乱数を生成
    b := make([]byte, digit)
    if _, err := rand.Read(b); err != nil {
        return "", errors.New("error failed generate state")
    }

    // letters からランダムに取り出して文字列を生成
    var state string
    for _, v := range b {
        // index が letters の長さに収まるように調整
        state += string(letters[int(v)%len(letters)])
    }
    return state, nil
}

func GetAuthCode() echo.HandlerFunc {
	return func(c echo.Context) error {
		config := gyazo.GetConnect()
		state, err := stateGenerator(10)
		if err != nil {
			logrus.Error(err)
		}
		url := config.AuthCodeURL(state)
		res := &AuthResponse{url, state}
		return c.JSON(http.StatusOK, res)
	}
}

func GetToken() echo.HandlerFunc {
	return func(c echo.Context) error {
		// アクセストークンの取得
		t := new(TokenRequest)
		if err := c.Bind(t); err != nil {
			logrus.Error("Error Request: ", err)
			return c.JSON(http.StatusBadRequest, err)
		}
		config := gyazo.GetConnect()

		token, err := config.Exchange(context.Background(), t.Code)
		if err != nil {
			logrus.Error("Error Request: ", err)
			return c.JSON(http.StatusBadRequest, err)
		}
		if !token.Valid() {
			logrus.Error("Invalid token: ", err)
			return c.JSON(http.StatusBadRequest, err)
		}

		// ユーザの取得
		client, err := gyazo.NewClient(token.AccessToken)
		if err != nil {
			logrus.Error(err)
		}
		gyazo_user, err := gyazo.GetUser(*client)
		if err != nil {
			logrus.Error(err)
		}
		// FirebaseでのJWT token取得
		authClient := c.Get("firebase").(*auth.Client)
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		user := models.User{}
		// DBにGyazoのUIDがなければ、FirebaseとDBに登録する
		err = dbs.DB.Table("users").Where(models.User{GyazoUID: gyazo_user.User.UID}).First(&user).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			params := (&auth.UserToCreate{}).
					Email(gyazo_user.User.Email).
					EmailVerified(false).
					Password(gyazo_user.User.UID).
					Disabled(false)
			u, err := authClient.CreateUser(context.Background(), params)
			if err != nil {
				logrus.Fatal("error creating user: ", err)
			}
			logrus.Info("create user: ", u)
			// db追加
			user = models.User{
				UID: u.UID, 
				AccessToken: token.AccessToken,
				Email: gyazo_user.User.Email,
				GyazoUID: gyazo_user.User.UID,
			}
			dbs.DB.Create(&user)
		}
		// DBには存在するが、メールアドレスが変わっている場合修正する
		if user.Email != gyazo_user.User.Email {
			user.Email = gyazo_user.User.Email
			dbs.DB.Save(&user)
		}
		jwt_token, _ := authClient.CustomToken(context.Background(), user.UID)
		// res := &TokenResponse{token.AccessToken}
		return c.JSON(http.StatusOK, jwt_token)
	}
}
