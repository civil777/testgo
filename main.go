package main

import (
	"html/template"
	"github.com/labstack/echo"
	"io"
	"net/http"
)

type TemplateRenderer struct {
	templates *template.Template
}
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	e := echo.New()
	e.Static("/assets", "assets")
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("*.html")),
	}
	e.Renderer = renderer

	// Named route "foobar"
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "landing.html", map[string]interface{}{
			"name": "Dolly!",
		})
	}).Name = "foobar"

	e.Logger.Fatal(e.Start(":8000"))
}