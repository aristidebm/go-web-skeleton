package main

import (
	"html/template"
	"io"
	"time"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data any) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type TemplateData struct {
	Form        any
	CurrentYear int
}

func NewTemplateData() TemplateData {
	return TemplateData{
		CurrentYear: time.Now().UTC().Year(),
	}
}

func GetTemplate(page string) (templ *template.Template, name string, err error) {
	name = "base"
	templ, err = template.New(name).ParseFiles(
		"ui/html/base.html",
		"ui/html/partials/navbar.html",
		page,
	)
	if err != nil {
		return nil, "", err
	}
	return
}
