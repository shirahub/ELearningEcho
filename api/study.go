package api

import (
	"E-LearningEcho/model"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func LetsStudy(c echo.Context) error {

	cookie, err := c.Cookie("token")
	if err != nil {
		return err
	}
	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)

	tknStr := cookie.Value

	// Initialize a new instance of `Claims`
	claims := &model.Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return c.String(http.StatusUnauthorized, "JWT invalid")

		}
		return c.String(http.StatusBadGateway, "Status Bad Gateaway")

	}
	if !tkn.Valid {
		return c.String(http.StatusUnauthorized, "JWT invalid")

	}
	data := model.M{"name": claims.Name}
	return c.Render(http.StatusOK, "study.html", data)

}
