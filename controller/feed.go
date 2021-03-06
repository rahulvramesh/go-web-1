package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

type Feed struct{}

func (ctrl Feed) GetFeed(c echo.Context) error {
	var user struct {
		Name string `json:"name"`
	}
	user.Name = "Peter"
	return c.JSON(http.StatusOK, user)
}
