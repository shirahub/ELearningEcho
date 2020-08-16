package api

import (
	"E-LearningEcho/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func MakePractice(c echo.Context) error {

	data := model.M{"message": "buat latihan"}
	return c.Render(http.StatusOK, "makepractice.html", data)
}

func SavePractice(c echo.Context) error {
	tingkat := c.FormValue("tingkat")
	kelas := c.FormValue("kelas")
	mapel := c.FormValue("mapel")
	tema := c.FormValue("tema")

	fmt.Println(tingkat, kelas, mapel, tema)

	c.Request().ParseForm()
	questions := c.Request().Form["question"]
	answers := c.Request().Form["answer"]
	choices1 := c.Request().Form["choices1"]
	choices2 := c.Request().Form["choices2"]
	choices3 := c.Request().Form["choices3"]

	fmt.Println("questions", questions)
	fmt.Println("answers", answers)
	fmt.Println("choices1", choices1)
	fmt.Println("choices2", choices2)
	fmt.Println("choices3", choices3)

	data := model.M{"message": "buat latihan success"}
	return c.Render(http.StatusOK, "user.html", data)
}
