package api

import (
	"E-LearningEcho/auth"
	"E-LearningEcho/model"
	"net/http"

	"github.com/labstack/echo"
)

func LetsStudy(c echo.Context) error {

	_, err := auth.IsSessionActive(c)
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": "User is not logged in"})
	}

	return c.Render(http.StatusOK, "study.html", model.M{"message": "User is on"})

}
