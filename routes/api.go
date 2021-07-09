package routes

import (
	"github.com/Wisteria30/J-analyzer/web/api"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	g := e.Group("/v1")
	{
		g.GET("/authorize", api.GetAuthCode())
		g.GET("/callback", api.GetToken())
		g.GET("/images", api.GetImages())
		g.GET("/hello", api.GetHello())
		// g.GET("/auth", api.GetAuthURL())
		// g.GET("/access-token", api.GetAccessToken())
	}
}
