package routes

import (
	"github.com/ImanAski/janotan-api/controllers"
	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Echo) {
	e.POST("/auth/login", controllers.Login)
	e.POST("/auth/register", controllers.Register)
}
