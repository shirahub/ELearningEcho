package api

import (
	"E-LearningEcho/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

var IsLoggedIn = []byte("secret")

func LoginPage(c echo.Context) error {
	data := model.M{"message": "This is the Login Page"}
	return c.Render(http.StatusOK, "login.html", data)
}

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
	getUserData := db.QueryRow("SELECT id_user, username, fullname FROM user WHERE username=?", user.Username).Scan(&userData.UserId, &userData.Username, &userData.Fullname)
	if getUserData != nil {
		// kalau ga ketemu / nil, kirim message NOT FOUND
		log.Fatal(err)
		data := model.M{"message": "NOT FOUND"}

		return c.Render(http.StatusOK, "homepage.html", data)
	}
	fmt.Println(userData)

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	// This is the information which frontend can use
	// The backend can also decode the token and get admin etc.
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = userData.Fullname
	claims["admin"] = true
	claims["userid"] = userData.UserId
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	// Generate encoded token and send it as response.
	// The signing string should be secret (a generated UUID          works too)
	t, err := token.SignedString(IsLoggedIn)
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.HttpOnly = true
	cookie.Name = "token"
	cookie.Value = t
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)

	pslist := ShowAllPractice()

	data := model.M{"message": "Login Success", "name": userData.Fullname, "pslist": pslist}
	return c.Render(http.StatusOK, "user.html", data)

	// return c.JSON(http.StatusOK, map[string]string{
	// 	"token": t,
	// })
}

func ShowAllPractice() string {
	db, err := sql.Open("mysql", "root:120625@/elearning")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var pslist []model.PaketSoal
	getPaketSoal, err := db.Query("SELECT id_paketsoal, tingkat, kelas, mapel, tema FROM paketsoal")
	if err != nil {
		panic(err)
	}
	defer getPaketSoal.Close()
	for getPaketSoal.Next() {
		ps := model.PaketSoal{}
		err = getPaketSoal.Scan(&ps.Id_PaketSoal, &ps.Tingkat, &ps.Kelas, &ps.Mapel, &ps.Tema)
		if err != nil {
			panic(err)
		}
		fmt.Println(ps)
		pslist = append(pslist, ps)
	}
	err = getPaketSoal.Err()
	if err != nil {
		panic(err)
	}

	fmt.Println(pslist)
	psjs, err := json.Marshal(pslist)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(psjs))
	return string(psjs)
}
