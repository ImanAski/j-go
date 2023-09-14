package routes

import (
	"github.com/ImanAski/janotan-api/controllers"
	"github.com/labstack/echo/v4"
)

func BookRoutes(e *echo.Echo) {
	e.GET("/books", controllers.GetBooks)
}
