package api

import (
	"E-LearningEcho/auth"
	"E-LearningEcho/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func LoginPage(c echo.Context) error {
	data := model.M{"message": ""}
	return c.Render(http.StatusOK, "login.html", data)
}

func UserLogin(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
	}

	db, err := sql.Open("mysql", "root:120625@/elearning")
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
	}
	defer db.Close()

	var userData model.UserData
	getUserDataErr := db.QueryRow("SELECT id_user, username, fullname FROM user WHERE username=? and password=?", user.Username, user.Password).Scan(&userData.UserId, &userData.Username, &userData.Fullname)
	if getUserDataErr != nil {
		return c.Render(http.StatusOK, "login.html", model.M{"message": "Username / Password invalid"})
	}

	t, err := generateToken(userData)
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
	}

	setTokenInCookie(c, t)

	pslist, err := ShowAllPractice()
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
	}

	userpslist, err := ShowAllPracticeByUser(userData.UserId)
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
	}
	userhistorylist, err := ShowUserHistory(userData.UserId)
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
	}

	data := model.M{"name": userData.Fullname, "pslist": pslist, "userpslist": userpslist, "userhistorylist": userhistorylist}
	return c.Render(http.StatusOK, "user.html", data)
}

func UserMenu(c echo.Context) error {
	_, err := auth.IsSessionActive(c)
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": "User is not logged in"})
	}

	username, id, err := GetUsernameAndUserIdFromToken(c)
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
	}
	pslist, err := ShowAllPractice()
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
	}
	userpslist, err := ShowAllPracticeByUser(id)
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
	}
	userhistorylist, err := ShowUserHistory(id)
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
	}

	data := model.M{"name": username, "pslist": pslist, "userpslist": userpslist, "userhistorylist": userhistorylist}
	return c.Render(http.StatusOK, "user.html", data)
}

func generateToken(userData model.UserData) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = userData.Fullname
	claims["admin"] = true
	claims["userid"] = userData.UserId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	t, err := token.SignedString([]byte("secret"))
	return t, err
}

func setTokenInCookie(c echo.Context, token string) {
	cookie := new(http.Cookie)
	cookie.HttpOnly = true
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
}

func ShowAllPractice() (string, error) {
	db, err := sql.Open("mysql", "root:120625@/elearning")
	if err != nil {
		return "", err
	}
	defer db.Close()

	var pslist []model.PaketSoal
	getPaketSoal, err := db.Query("SELECT id_paketsoal, tingkat, kelas, mapel, tema FROM paketsoal")
	if err != nil {
		return "", err
	}
	defer getPaketSoal.Close()
	for getPaketSoal.Next() {
		ps := model.PaketSoal{}
		err = getPaketSoal.Scan(&ps.Id_PaketSoal, &ps.Tingkat, &ps.Kelas, &ps.Mapel, &ps.Tema)
		if err != nil {
			return "", err
		}
		pslist = append(pslist, ps)
	}
	err = getPaketSoal.Err()
	if err != nil {
		return "", err
	}

	psjs, err := json.Marshal(pslist)
	if err != nil {
		return "", err
	}
	return string(psjs), nil
}

func ShowAllPracticeByUser(id int) (string, error) {

	db, err := sql.Open("mysql", "root:120625@/elearning")
	if err != nil {
		return "", err
	}
	defer db.Close()

	var pslist []model.PaketSoal
	getPaketSoal, err := db.Query("SELECT id_paketsoal, tingkat, kelas, mapel, tema FROM paketsoal where id_user=?", id)
	if err != nil {
		return "", err
	}
	defer getPaketSoal.Close()
	for getPaketSoal.Next() {
		ps := model.PaketSoal{}
		err = getPaketSoal.Scan(&ps.Id_PaketSoal, &ps.Tingkat, &ps.Kelas, &ps.Mapel, &ps.Tema)
		if err != nil {
			return "", err
		}
		pslist = append(pslist, ps)
	}
	err = getPaketSoal.Err()
	if err != nil {
		return "", err
	}

	psjs, err := json.Marshal(pslist)
	if err != nil {
		return "", err
	}
	return string(psjs), nil
}

func ShowUserHistory(id int) (string, error) {

	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println(err)
	}

	db, err := sql.Open("mysql", "root:120625@/elearning?parseTime=true")
	if err != nil {
		return "", err
	}
	defer db.Close()

	var uhlist []model.UserHistory
	getUserHistory, err := db.Query("SELECT id_paketsoal, nilai, waktu FROM userhistory where id_user=?", id)
	if err != nil {
		return "", err
	}
	defer getUserHistory.Close()
	for getUserHistory.Next() {
		uh := model.UserHistory{}
		err = getUserHistory.Scan(&uh.PaketId, &uh.Nilai, &uh.Waktu)
		if err != nil {
			return "", err
		}
		uhlist = append(uhlist, uh)
	}
	err = getUserHistory.Err()
	if err != nil {
		return "", err
	}

	var pslist []model.PaketSoal
	for i := 0; i < len(uhlist); i++ {
		getPaketSoal, err := db.Query("SELECT id_paketsoal, tingkat, kelas, mapel, tema FROM paketsoal where id_paketsoal=?", uhlist[i].PaketId)
		if err != nil {
			return "", err
		}
		defer getPaketSoal.Close()
		for getPaketSoal.Next() {
			ps := model.PaketSoal{}
			err = getPaketSoal.Scan(&ps.Id_PaketSoal, &ps.Tingkat, &ps.Kelas, &ps.Mapel, &ps.Tema)
			if err != nil {
				return "", err
			}
			pslist = append(pslist, ps)
		}
		err = getPaketSoal.Err()
		if err != nil {
			return "", err
		}

		uhlist[i].NamaPaket = pslist[i].Tingkat + "-" + pslist[i].Kelas + "-" + pslist[i].Mapel + "-" + pslist[i].Tema

		uhlist[i].Waktu = uhlist[i].Waktu.In(loc)
		uhlist[i].Waktustring = uhlist[i].Waktu.Format("2006-01-02 15:04:05")
	}

	uhjs, err := json.Marshal(uhlist)
	if err != nil {
		return "", err
	}
	return string(uhjs), nil
}
