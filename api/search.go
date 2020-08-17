package api

import (
	"E-LearningEcho/model"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

func SearchPractice(c echo.Context) error {
	db, err := sql.Open("mysql", "root:120625@/elearning")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var pslist []model.PaketSoal
	getPaketSoal, err := db.Query("SELECT * FROM paketsoal")
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
		pslist = append(pslist, ps)
	}
	err = getPaketSoal.Err()
	if err != nil {
		panic(err)
	}

	psjs, err := json.Marshal(pslist)
	if err != nil {
		panic(err)
	}
	data := model.M{"message": string(psjs)}

	return c.Render(http.StatusOK, "user.html", data)
}
