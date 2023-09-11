package controllers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ImanAski/janotan-api/database"
	"github.com/ImanAski/janotan-api/models"
	"github.com/ImanAski/janotan-api/responses"
	"github.com/ImanAski/janotan-api/utils"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	if len(email) == 0 || len(password) == 0 {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "email or password is missing",
			Data:    nil,
		})
	}
	godotenv.Load(".env")
	user := models.User{
		Email: email,
	}

	result := database.DB.Db.First(&user)

	if result.Error != nil {
		log.Println("There was an error", result.Error)
	}

	if user.Email != email || !utils.CheckPasswordHash(password, user.Password) {
		userResponse := responses.UserResponse{
			Status:  http.StatusUnauthorized,
			Message: "the email or password was incorrect",
			Data:    nil,
		}
		return c.JSON(http.StatusUnauthorized, userResponse)
	}
	// claims := &models.JwtCustomClaims{
	// 	user.Name + " " + user.LastName,
	// 	false,
	// 	jwt.RegisteredClaims{
	// 		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
	// 	},
	// }

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Name + " " + user.LastName
	claims["admin"] = user.Admin
	claims["exp"] = time.Now().UTC().Add(time.Hour).Unix()

	t, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return err
	}
	userResponse := responses.UserResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data: &echo.Map{
			"token": t,
		},
	}
	return c.JSON(http.StatusOK, userResponse)
}

func Register(c echo.Context) error {

	userResponse := responses.UserResponse{
		Status:  http.StatusOK,
		Message: "User Registered",
		Data:    nil,
	}
	name := c.FormValue("name")
	lastName := c.FormValue("last_name")
	phoneNumber := c.FormValue("phone_number")
	email := c.FormValue("email")
	birthdate := c.FormValue("birthdate")
	password := c.FormValue("password")

	passwordHash, err := utils.PasswordHash(password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "there was something wrong with server",
			Data:    nil,
		})
	}

	user := models.User{
		Name:        name,
		LastName:    lastName,
		Email:       email,
		Password:    passwordHash,
		Birthdate:   birthdate,
		PhoneNumber: phoneNumber,
		Balance:     0,
		Admin:       false,
	}

	result := database.DB.Db.Create(&user)
	if result.Error != nil {
		log.Fatal("There was a problem")
		userResponse := responses.UserResponse{
			Status:  http.StatusUnauthorized,
			Message: "There was a problem inserting user.",
			Data: &echo.Map{
				"error": result.Error,
			},
		}
		return c.JSON(http.StatusUnauthorized, userResponse)
	}
	return c.JSON(http.StatusOK, userResponse)
}
