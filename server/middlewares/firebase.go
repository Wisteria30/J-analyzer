package middlewares

import (
	"context"
	"os"

	firebase "firebase.google.com/go/v4"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

func Firebase() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			opt := option.WithCredentialsFile(os.Getenv("KEY_JSON_PATH"))
			config := &firebase.Config{ProjectID: os.Getenv("PROJECT_ID")}
			app, err := firebase.NewApp(context.Background(), config, opt)
			if err != nil {
				logrus.Fatalf("Error initializing firebase: %v\n", err)
			}

			auth, err := app.Auth(context.Background())
			if err != nil {
				return err
			}

			c.Set("firebase", auth)
 
			if err := next(c); err != nil {
				return err
			}

			return nil
		}
	}
}