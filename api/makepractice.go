package api

import (
	"E-LearningEcho/model"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func MakePractice(c echo.Context) error {

	data := model.M{"message": "buat latihan"}
	return c.Render(http.StatusOK, "makepractice.html", data)
}

func SavePractice(c echo.Context) error {

	name, idUser, err := GetUsernameAndUserIdFromToken(c)
	if err != nil {
		fmt.Println("error")
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
	fmt.Println(name, tingkat, kelas, mapel, tema, idUser)
	fmt.Println("questions", questions)
	fmt.Println("answers", answers)
	fmt.Println("choices1", choices1)
	fmt.Println("choices2", choices2)
	fmt.Println("choices3", choices3)

	db, err := sql.Open("mysql", "root:120625@/elearning")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO paketsoal(tingkat, kelas, mapel, tema, id_user) VALUES(?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(tingkat, kelas, mapel, tema, idUser)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
	lastPaketSoalId := lastId

	for i := 0; i < len(questions); i++ {
		stmt, err := db.Prepare("INSERT INTO soal(id_paketsoal, pertanyaan, jawaban, pilihan1, pilihan2, pilihan3) VALUES(?,?,?,?,?,?)")
		if err != nil {
			log.Fatal(err)
		}
		res, err := stmt.Exec(lastPaketSoalId, questions[i], answers[i], choices1[i], choices2[i], choices3[i])
		if err != nil {
			log.Fatal(err)
		}
		lastId, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
	}

	data := model.M{"message": "buat latihan success"}
	return c.Render(http.StatusOK, "user.html", data)
}
