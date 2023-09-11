package main

import (
	"log"

	"github.com/ImanAski/janotan-api/database"
	"github.com/ImanAski/janotan-api/routes"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// Run Database
	database.ConnectDb()

	// Run Routes

	// * User Routes
	routes.UserRoute(e)

	// * Auth Routes
	routes.AuthRoutes(e)

	//middleware

	// Run Application
	log.Fatal(e.Start(":3000"))
}
