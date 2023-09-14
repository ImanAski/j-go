package controllers

import (
	"net/http"
	"strconv"

	"github.com/ImanAski/janotan-api/database"
	"github.com/ImanAski/janotan-api/models"
	"github.com/ImanAski/janotan-api/responses"
	"github.com/labstack/echo/v4"
)

func GetCourses(c echo.Context) error {
	courses := []models.Course{}

	result := database.DB.Db.Find(&courses)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "there was an error retrieving courses",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, responses.UserResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data: &echo.Map{
			"courses": courses,
		},
	})
}

func CreateCourse(c echo.Context) error {
	name := c.FormValue("name")
	description := c.FormValue("description")
	price, err := strconv.ParseFloat(c.FormValue("price"), 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "there was a problem parsing price value",
			Data:    nil,
		})
	}
	sellingPrice, err := strconv.ParseFloat(c.FormValue("selling_price"), 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "there was a problem parsing selling price value",
			Data:    nil,
		})
	}

	course := models.Course{
		Name:         name,
		Description:  description,
		Price:        price,
		SellingPrice: sellingPrice,
	}

	result := database.DB.Db.Create(&course)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "there was a error inserting the course",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, responses.UserResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data: &echo.Map{
			"course": course,
		},
	})
}
