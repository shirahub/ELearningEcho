package api

import (
	"E-LearningEcho/auth"
	"E-LearningEcho/model"
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
)

func MakePractice(c echo.Context) error {
	_, err := auth.IsSessionActive(c)
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": "User is not logged in"})
	}

	data := model.M{"message": "buat latihan"}
	return c.Render(http.StatusOK, "makepractice.html", data)
}

func SavePractice(c echo.Context) error {

	_, err := auth.IsSessionActive(c)
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": "User is not logged in"})
	}

	_, idUser, err := GetUsernameAndUserIdFromToken(c)
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
	}

	tingkat := c.FormValue("tingkat")
	kelas := c.FormValue("kelas")
	mapel := c.FormValue("mapel")
	tema := c.FormValue("tema")

	c.Request().ParseForm()
	questions := c.Request().Form["question"]
	answers := c.Request().Form["answer"]
	choices1 := c.Request().Form["choices1"]
	choices2 := c.Request().Form["choices2"]
	choices3 := c.Request().Form["choices3"]

	// if tingkat =="" || kelas =="" || mapel==""|| tema =="" {
	// 	return c.Render(http.StatusOK, "error.html", model.M{"message": "Please fill in the "})
	// }

	db, err := sql.Open("mysql", "root:120625@/elearning")
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO paketsoal(tingkat, kelas, mapel, tema, id_user) VALUES(?,?,?,?,?)")
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
	}
	res, err := stmt.Exec(tingkat, kelas, mapel, tema, idUser)
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
	}

	lastPaketSoalId := lastId

	for i := 0; i < len(questions); i++ {
		stmt, err := db.Prepare("INSERT INTO soal(id_paketsoal, pertanyaan, jawaban, pilihan1, pilihan2, pilihan3) VALUES(?,?,?,?,?,?)")
		if err != nil {
			return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
		}
		_, err = stmt.Exec(lastPaketSoalId, questions[i], answers[i], choices1[i], choices2[i], choices3[i])
		if err != nil {
			return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
		}
	}

	data := model.M{"message": "buat latihan success"}
	return c.Render(http.StatusOK, "makesuccess.html", data)
}

func SaveEdittedPractice(c echo.Context) error {

	_, err := auth.IsSessionActive(c)
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": "User is not logged in"})
	}

	// _, idUser, err := GetUsernameAndUserIdFromToken(c)
	// if err != nil {
	// 	return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
	// }
	id := c.Param("id")
	tingkat := c.FormValue("tingkat")
	kelas := c.FormValue("kelas")
	mapel := c.FormValue("mapel")
	tema := c.FormValue("tema")

	c.Request().ParseForm()
	questions := c.Request().Form["question"]
	answers := c.Request().Form["answer"]
	choices1 := c.Request().Form["choices1"]
	choices2 := c.Request().Form["choices2"]
	choices3 := c.Request().Form["choices3"]

	// fmt.Println(tingkat, kelas, mapel, tema)
	// fmt.Println(questions)
	// fmt.Println(answers)
	// fmt.Println(choices1)
	// fmt.Println(choices2)
	// fmt.Println(choices3)

	db, err := sql.Open("mysql", "root:120625@/elearning")
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
	}
	defer db.Close()

	updatePaketSoal, err := db.Prepare("UPDATE paketsoal SET tingkat=?, kelas=?, mapel=?, tema=? WHERE id_paketsoal=?")
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
	}
	updatePaketSoal.Exec(tingkat, kelas, mapel, tema, id)

	delSoal, err := db.Prepare("DELETE FROM soal WHERE id_paketsoal=?")
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
	}
	delSoal.Exec(id)

	for i := 0; i < len(questions); i++ {
		stmt, err := db.Prepare("INSERT INTO soal(id_paketsoal, pertanyaan, jawaban, pilihan1, pilihan2, pilihan3) VALUES(?,?,?,?,?,?)")
		if err != nil {
			return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
		}
		stmt.Exec(id, questions[i], answers[i], choices1[i], choices2[i], choices3[i])
	}

	data := model.M{"message": "Edit latihan success"}
	return c.Render(http.StatusOK, "makesuccess.html", data)
}
