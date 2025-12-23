package views

import (
	"embed"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

//go:embed *.html
var viewFS embed.FS

type Renderer struct {
	templates *template.Template
}

func (r *Renderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	return r.templates.ExecuteTemplate(w, name, data)
}

func NewRenderer() *Renderer {
	funcMap := template.FuncMap{
		"plus":  func(a, b int) int { return a + b },
		"minus": func(a, b int) int { return a - b },
	}

	t := template.Must(
		template.New("").Funcs(funcMap).ParseFS(viewFS, "*.html"),
	)

	return &Renderer{
		templates: t,
	}
}
