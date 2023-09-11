package controllers

import (
	"net/http"

	"github.com/ImanAski/janotan-api/responses"
	"github.com/labstack/echo"
)

func CreateUser(c echo.Context) error {

	r := responses.UserResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data: &echo.Map{
			"test": "success",
		},
	}
	return c.JSON(http.StatusOK, r)
}
