package api

import (
	"E-LearningEcho/model"
	"net/http"

	"github.com/labstack/echo"
)

func HomePage(c echo.Context) error {
	data := model.M{"message": "func HomePage"}
	return c.Render(http.StatusOK, "login.html", data)
}
