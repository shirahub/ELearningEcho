package api

import (
	"E-LearningEcho/model"
	"net/http"

	"github.com/labstack/echo"
)

func HomePage(c echo.Context) error {
	data := model.M{"message": "This is the homepage"}
	return c.Render(http.StatusOK, "homepage.html", data)
}

//kirim data tapi ga dipake
