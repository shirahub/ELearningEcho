package api

import (
	"E-LearningEcho/model"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func UserLogin(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	// Open up our database connection.
	db, err := sql.Open("mysql", "root:120625@/elearning")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	defer db.Close()

	//single row query lalu langsung parse ke model.UserData
	// referensi: http://go-database-sql.org/retrieving.html
	var userData model.UserData
	getUserData := db.QueryRow("SELECT username, fullname FROM user WHERE username=?", user.Username).Scan(&userData.Username, &userData.Fullname)
	if getUserData != nil {
		// kalau ga ketemu / nil, kirim message NOT FOUND
		log.Fatal(err)
		data := model.M{"message": "NOT FOUND"}

		return c.Render(http.StatusOK, "user.html", data)
	}
	fmt.Println(userData)

	// kalau ketemu, kirim userData yang didapatkan
	data := model.M{"message": userData}
	return c.Render(http.StatusOK, "user.html", data)
}
