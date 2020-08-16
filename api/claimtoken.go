package api

import (
	"E-LearningEcho/model"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func GetUsernameAndUserIdFromToken(c echo.Context) (string, int, error) {
	cookie, err := c.Cookie("token")
	if err != nil {
		return "token not found in cookie", 0, err
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
			return 0, err

		}
		return 0, err

	}
	if !tkn.Valid {
		return 0, err

	}
	return claims.UserId, nil
}
