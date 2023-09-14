package routes

import (
	"github.com/ImanAski/janotan-api/controllers"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	r := e.Group("/user")

	r.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("secret"),
	}))
	r.GET("/", controllers.GetUser)
}
