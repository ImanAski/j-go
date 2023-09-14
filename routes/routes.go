package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ImanAski/janotan-api/responses"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.GET("/routes", func(c echo.Context) error {
		log.Println(e.Routes())
		data, err := json.MarshalIndent(e.Routes(), "", "  ")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "there was something wrong with the server",
				Data:    nil,
			})
		}

		return c.JSON(http.StatusOK, responses.UserResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data: &echo.Map{
				"routes": data,
			},
		})
	})
}
