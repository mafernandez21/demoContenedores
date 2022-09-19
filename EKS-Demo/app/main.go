package main

import (
	"html/template"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

const fortuneFile = "./fortune.txt"

// TEMPLATES
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// FORTUNE TELLER
func readFortuneFile(fortuneFile string) ([]string, error) {
	content, err := os.ReadFile(fortuneFile)
	var fortunes []string = nil
	if err == nil {
		fortunes = strings.Split(string(content), "\n%\n")
		return fortunes, err
	}
	return fortunes, err
}

func findAndPrint(fortuneFile string) (string, error) {
	fortunes, err := readFortuneFile(fortuneFile)
	if err == nil {
		rand.Seed(time.Now().UTC().UnixNano())
		i := rand.Int() % len(fortunes)
		return fortunes[i], err
	}
	return "", err
}

// HANDLERS
func healthcheck(c echo.Context) error {
	return c.String(http.StatusOK, "It's Working...")
}

func fortune(c echo.Context) error {
	name := c.QueryParam("fname")
	fortune, _ := findAndPrint(fortuneFile)
	return c.Render(http.StatusOK, "fortune.html", map[string]interface{}{
		"name":       name,
		"fortunemsg": fortune,
	})
}

func main() {

	e := echo.New()

	//TEMPLATES
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e.Renderer = t

	//STATIC FILES
	e.Static("/css", "public/css")
	e.Static("/js", "public/js")
	e.File("/", "public/index.html")

	//ROUTING
	e.GET("/healthcheck", healthcheck)
	e.GET("/fortune", fortune)

	//SERVER START
	e.Logger.Fatal(e.Start(":3000"))
}
