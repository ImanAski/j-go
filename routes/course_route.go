package routes

import (
	"github.com/ImanAski/janotan-api/controllers"
	"github.com/labstack/echo/v4"
)

func CourseRoute(e *echo.Echo) {
	r := e.Group("/course")

	r.GET("", controllers.GetCourses)
	r.POST("", controllers.CreateCourse)
}
