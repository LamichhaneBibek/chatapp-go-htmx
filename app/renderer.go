package main

import (
	"bytes"
	"fmt"
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
)

type HTMLRenderer struct {
	templates *template.Template
}

func NewHtmlRenderer(path string) *HTMLRenderer {
	fmt.Println("it accept .html extension")
	return &HTMLRenderer{
		templates: template.Must(template.ParseGlob(path + "/*.html")),
	}
}

func (r *HTMLRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return r.templates.ExecuteTemplate(w, name, data)
}

func (r *HTMLRenderer) RenderToString(name string, data interface{}) (string, error) {
	var buf []byte
	w := bytes.NewBuffer(buf)

	err := r.Render(w, name, data, nil)
	if err != nil {
		return "", err
	}

	return w.String(), nil
}
