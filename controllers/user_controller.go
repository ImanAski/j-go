package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {

	return c.JSON(http.StatusOK, echo.Map{
		"token": "hello",
	})
}
