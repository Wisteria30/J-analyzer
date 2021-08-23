package routes

import (
	"github.com/Wisteria30/J-analyzer/middlewares"
	"github.com/Wisteria30/J-analyzer/web/api"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	g := e.Group("/v1")
	{
		g.GET("/authorize", api.GetAuthCode())
		g.GET("/token", api.GetToken())
		g.GET("/images", api.GetList(), middlewares.FirebaseGuard())
		g.GET("/images/:id", api.GetImage(), middlewares.FirebaseGuard())
		g.GET("/analyze", api.AnalyzeImage())
	}
}
