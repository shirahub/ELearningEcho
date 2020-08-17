package api

import (
	"E-LearningEcho/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func ShowPractice(c echo.Context) error {
	id := c.Param("id")
	tingkat := c.Param("tingkat")
	kelas := c.Param("kelas")
	mapel := c.Param("mapel")
	tema := c.Param("tema")

	db, err := sql.Open("mysql", "root:120625@/elearning")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var soallist []model.Soal
	getSoal, err := db.Query("SELECT id_soal, pertanyaan, jawaban, pilihan1, pilihan2, pilihan3 FROM soal WHERE id_paketsoal=?", id)
	if err != nil {
		panic(err)
	}
	defer getSoal.Close()
	for getSoal.Next() {
		so := model.Soal{}
		rand.Seed(time.Now().UnixNano())
		p := rand.Perm(4)

		err = getSoal.Scan(&so.Id_Soal, &so.Pertanyaan, &so.Pilihan[p[0]], &so.Pilihan[p[1]], &so.Pilihan[p[2]], &so.Pilihan[p[3]])
		if err != nil {
			panic(err)
		}
		// sojs, err := json.Marshal(so)
		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Println(string(sojs))

		// fmt.Println(so)
		soallist = append(soallist, so)
	}
	err = getSoal.Err()
	if err != nil {
		panic(err)
	}

	fmt.Println(soallist)
	sojs, err := json.Marshal(soallist)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(sojs))
	data := model.M{"message": string(sojs), "tingkat": tingkat, "kelas": kelas, "mapel": mapel, "tema": tema, "id": id}

	return c.Render(http.StatusOK, "dopractice.html", data)
	// return c.String(http.StatusOK, "asdasd")
}

func GetAnswers(c echo.Context) error {
	id := c.Param("id")
	tingkat := c.Param("tingkat")
	kelas := c.Param("kelas")
	mapel := c.Param("mapel")
	tema := c.Param("tema")

	db, err := sql.Open("mysql", "root:120625@/elearning")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var answercount int
	var answerlist []string
	getAnswer, err := db.Query("SELECT jawaban FROM soal WHERE id_paketsoal=?", id)
	if err != nil {
		panic(err)
	}
	defer getAnswer.Close()
	for getAnswer.Next() {
		var answer string
		err = getAnswer.Scan(&answer)
		if err != nil {
			panic(err)
		}
		answercount++
		answerlist = append(answerlist, answer)
	}

	var correct int
	for i := 0; i < answercount; i++ {
		stringparam := "answerno" + strconv.Itoa(i)
		answer := c.FormValue(stringparam)
		// fmt.Println("answer dari user", answer)
		// fmt.Println(answerlist[i])

		if answer == answerlist[i] {
			correct++
		}
	}

	score := (100 / answercount) * correct

	data := model.M{"tingkat": tingkat, "kelas": kelas, "mapel": mapel, "tema": tema, "id": id, "result": score}

	return c.Render(http.StatusOK, "result.html", data)
}
