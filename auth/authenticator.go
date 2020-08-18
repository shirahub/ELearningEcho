package auth

import (
	"E-LearningEcho/model"
	"database/sql"

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

func IsUserTheMaker(id_user int, id_paketsoal int) (bool, error) {
	db, err := sql.Open("mysql", "root:120625@/elearning")
	if err != nil {
		return false, err
	}
	defer db.Close()

	var id_maker int
	getMaker := db.QueryRow("SELECT id_user FROM paketsoal WHERE id_paketsoal=?", id_paketsoal).Scan(&id_maker)
	if getMaker != nil {
		return false, nil
	}

	if id_user != id_maker {
		return false, nil
	}

	return true, nil

}
