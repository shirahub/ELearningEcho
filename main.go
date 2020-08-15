package main

import (
	api "E-LearningEcho/api"
	"io"
	"log"
	"os"
	"text/template"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// HTML RENDERER CONFIG start ============================================

type Renderer struct {
	template *template.Template
	debug    bool
	location string
}

func NewRenderer(location string, debug bool) *Renderer {
	tpl := new(Renderer)
	tpl.location = location
	tpl.debug = debug
	tpl.ReloadTemplates()
	return tpl
}

func (t *Renderer) ReloadTemplates() {
	t.template = template.Must(template.ParseGlob(t.location))
}

func (t *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if t.debug {
		t.ReloadTemplates()
	}

	return t.template.ExecuteTemplate(w, name, data)
}

// HTML RENDERER CONFIG end ============================================

// JWT CONFIG start ============================================
// var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
// 	SigningKey: []byte("secret"),
// })

// JWT CONFIG end ============================================

func main() {

	// Logging file path
	f, err := os.OpenFile("myLog", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	// Echo instance
	e := echo.New()

	// Register html Renderer
	e.Renderer = NewRenderer("view/*.html", true)

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}, method=${method}, uri=${uri}, status=${status}, error=${error}\n",
		Output: f,
	}))
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/home", api.HomePage) //access ini di browser
	e.GET("/loginpage", api.LoginPage)
	e.POST("/login", api.UserLogin) //ini diaccess oleh submit button
	e.GET("/search", api.SearchPractice)
	e.GET("/logout", api.UserLogout)
	e.GET("/dopractice/:id", api.ShowPractice)
	e.POST("/submitanswers", api.GetAnswers)

	r := e.Group("/user")
	r.GET("/letsstudy", api.LetsStudy)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))

}
