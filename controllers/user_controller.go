package controllers

import (
	"net/http"

	"github.com/ImanAski/janotan-api/database"
	"github.com/ImanAski/janotan-api/models"
	"github.com/ImanAski/janotan-api/responses"
	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {

	return c.JSON(http.StatusOK, echo.Map{
		"token": "hello",
	})
}

func GetUser(c echo.Context) error {
	email := c.FormValue("email")
	if len(email) == 0 {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "Email key is required",
			Data:    nil,
		})
	}

	user := models.User{
		Email: email,
	}

	database.DB.Db.Find(&user)
	if user.ID < 1 {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "the email could not be found",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, responses.UserResponse{
		Status:  http.StatusOK,
		Message: "user returned successfully",
		Data: &echo.Map{
			"user": user,
		},
	})
}
