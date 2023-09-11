package routes

import (
	"github.com/ImanAski/janotan-api/controllers"
	"github.com/labstack/echo"
)

func AuthRoutes(e *echo.Echo) {
	e.POST("/auth/login", controllers.Login)
}
