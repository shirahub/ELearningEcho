package auth

import (
	"E-LearningEcho/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func IsSessionActive(c echo.Context) (bool, error) {
	cookie, err := c.Cookie("token")
	if err != nil {
		return false, err
	}
	tknStr := cookie.Value
	claims := &model.Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, err

		}
		return false, err

	}
	if !tkn.Valid {
		return false, err

	}
	return true, nil
}
