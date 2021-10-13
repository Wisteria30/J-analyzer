package main

import (
	"github.com/Wisteria30/J-analyzer/middlewares"
	"github.com/Wisteria30/J-analyzer/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env")
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	e := echo.New()

	//Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middlewares.DatabaseService())
	e.Use(middlewares.Firebase())
	e.Use(middleware.CORS())

	routes.Init(e)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
