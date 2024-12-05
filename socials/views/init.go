package views

import (
	"html/template"
)

var (
	templates *template.Template
)

func Init(dir string) {
	templates = template.Must(template.ParseGlob("*.html"))
}
