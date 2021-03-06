package api

import (
	"E-LearningEcho/auth"
	"E-LearningEcho/model"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func UserLogout(c echo.Context) error {

	_, err := auth.IsSessionActive(c)
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": "User is not logged in"})
	}
	cookie := new(http.Cookie)
	cookie.HttpOnly = true
	cookie.Name = "token"
	cookie.Value = "empty"
	cookie.Expires = time.Now().Add(60 * time.Second)
	c.SetCookie(cookie)

	data := model.M{"message": "Logout Success"}
	return c.Render(http.StatusOK, "homepage.html", data)
}
