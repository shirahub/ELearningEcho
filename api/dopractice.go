package api

import (
	"E-LearningEcho/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func ShowPractice(c echo.Context) error {
	db, err := sql.Open("mysql", "root:120625@/elearning")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var soallist []model.Soal
	getSoal, err := db.Query("SELECT id_soal, pertanyaan, jawaban, pilihan1, pilihan2, pilihan3 FROM soal WHERE id_paketsoal=?", 1)
	if err != nil {
		panic(err)
	}
	defer getSoal.Close()
	for getSoal.Next() {
		so := model.Soal{}
		err = getSoal.Scan(&so.Id_Soal, &so.Pertanyaan, &so.Pilihan[0], &so.Pilihan[1], &so.Pilihan[2], &so.Pilihan[3])
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
	data := model.M{"message": string(sojs)}

	return c.Render(http.StatusOK, "dopractice.html", data)
	// return c.String(http.StatusOK, "asdasd")
}

func GetAnswers(c echo.Context) error {
	return c.String(http.StatusOK, "belum selesai")
}
