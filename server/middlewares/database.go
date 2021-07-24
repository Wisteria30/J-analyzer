package middlewares

import (
	"github.com/Wisteria30/J-analyzer/databases"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type DatabaseClient struct {
	DB *gorm.DB
}

func DatabaseService() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			session, _ := databases.Connect()
			d := DatabaseClient{DB: session}
			db, _ := d.DB.DB()
			defer db.Close()

			c.Set("dbs", &d)

			if err := next(c); err != nil {
				return err
			}

			return nil
		}
	}
}