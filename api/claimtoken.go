package api

import (
	"E-LearningEcho/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func GetUsernameAndUserIdFromToken(c echo.Context) (string, int, error) {
	cookie, err := c.Cookie("token")
	if err != nil {
		return "token not found in cookie", 0, err
	}

	tknStr := cookie.Value
	claims := &model.Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "JWT invalid", 0, err

		}
		return "Status Bad Gateaway", 0, err

	}
	if !tkn.Valid {
		return "JWT invalid", 0, err

	}
	return claims.Name, claims.UserId, nil
}

func GetUserIdFromToken(c echo.Context) (int, error) {
	cookie, err := c.Cookie("token")
	if err != nil {
		return 0, err
	}

	tknStr := cookie.Value

	claims := &model.Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return 0, err

		}
		return 0, err

	}
	if !tkn.Valid {
		return 0, err

	}
	return claims.UserId, nil
}
