package api

import (
	"E-LearningEcho/auth"
	"E-LearningEcho/model"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func EditPractice(c echo.Context) error {
	_, err := auth.IsSessionActive(c)
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
	}

	id_user, err := GetUserIdFromToken(c)
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
	}

	id := c.Param("id")
	idint, _ := strconv.Atoi(id)
	isUserTheMaker, err := auth.IsUserTheMaker(id_user, idint)
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
	}

	if isUserTheMaker == false {
		return c.Render(http.StatusOK, "error.html", model.M{"message": "You are not allowed to edit this practice"})
	}

	tingkat := c.Param("tingkat")
	kelas := c.Param("kelas")
	mapel := c.Param("mapel")
	tema := c.Param("tema")

	db, err := sql.Open("mysql", "root:120625@/elearning")
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
	}
	defer db.Close()

	var soallist []model.Soal
	getSoal, err := db.Query("SELECT id_soal, pertanyaan, jawaban, pilihan1, pilihan2, pilihan3 FROM soal WHERE id_paketsoal=?", id)
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
	}
	defer getSoal.Close()
	for getSoal.Next() {
		so := model.Soal{}

		err = getSoal.Scan(&so.Id_Soal, &so.Pertanyaan, &so.Pilihan[0], &so.Pilihan[1], &so.Pilihan[2], &so.Pilihan[3])
		if err != nil {
			return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
		}
		soallist = append(soallist, so)
	}
	err = getSoal.Err()
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
	}

	sojs, err := json.Marshal(soallist)
	if err != nil {
		return c.Render(http.StatusOK, "error.html", model.M{"message": err.Error()})
	}

	data := model.M{"message": string(sojs), "tingkat": tingkat, "kelas": kelas, "mapel": mapel, "tema": tema, "id": id}

	return c.Render(http.StatusOK, "editpractice.html", data)

}
