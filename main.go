package main

import (
	"log"

	"github.com/ImanAski/janotan-api/database"
	"github.com/ImanAski/janotan-api/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
	}))

	// Run Database
	database.ConnectDb()

	// Run Routes

	// * User Routes
	routes.UserRoute(e)

	// * Auth Routes
	routes.AuthRoutes(e)

	// * Course Routes
	routes.CourseRoute(e)

	// * Routes
	routes.Routes(e)

	//middleware

	// Run Application
	log.Fatal(e.Start(":3000"))
}
