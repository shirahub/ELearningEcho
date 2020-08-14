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

	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:120625@/elearning")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	var userData model.UserData
	getUserData := db.QueryRow("SELECT username, fullname FROM user WHERE username=?", user.Username).Scan(&userData.Username, &userData.Fullname)
	if getUserData != nil {
		//fmt.Println("ngga ketemu")
		log.Fatal(err)
		data := model.M{"message": "NOT FOUND"}

		return c.Render(http.StatusOK, "user.html", data)
	}
	fmt.Println(userData)

	data := model.M{"message": userData}
	return c.Render(http.StatusOK, "user.html", data)
}
