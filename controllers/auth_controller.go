package controllers

import (
	"log"
	"net/http"

	"github.com/ImanAski/janotan-api/database"
	"github.com/ImanAski/janotan-api/models"
	"github.com/ImanAski/janotan-api/responses"
	"github.com/ImanAski/janotan-api/utils"
	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user := models.User{
		Email:    email,
		Password: utils.PasswordHash(password),
	}

	result := database.DB.Db.First(&user)

	if result.Error != nil {
		log.Println("There was an error", result.Error)
	}

	if user.Email != email || user.Password != utils.PasswordHash(password) {
		userResponse := responses.UserResponse{
			Status:  http.StatusUnauthorized,
			Message: "The Email Or Password was Incorrect",
			Data:    nil,
		}
		return c.JSON(http.StatusUnauthorized, userResponse)
	}

	return echo.ErrUnauthorized
}
