package api

import (
	"E-LearningEcho/model"
	"database/sql"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func RegistPage(c echo.Context)error{
	data := model.M{"message": "This is the Login Page"}
	return c.Render(http.StatusOK, "regis.html", data)
}

func Regist(c echo.Context)error{
	users := new(model.UserRegist)
	if err:= c.Bind(users);
	err!=nil{
		return err
	}
	db,err :=sql.Open("mysql", "root:realmadrid@/elearning")
	if err!=nil{
		panic(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO user(username,password,fullname) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(users.Username,users.Password,users.Fullname)
	if err != nil {
		log.Fatal(err)
	}

	data := model.M{}
	return c.Render(http.StatusOK, "homepage.html", data)
}

