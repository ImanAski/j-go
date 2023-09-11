package routes

import (
	"github.com/ImanAski/janotan-api/controllers"
	"github.com/labstack/echo"
)

func UserRoute(e *echo.Echo) {
	e.POST("/user", controllers.CreateUser)
}
