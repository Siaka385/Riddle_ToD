package views

import (
	"net/http"
)

func Render(w http.ResponseWriter, r *http.Request, tName string, data any) {
	if err := templates.ExecuteTemplate(w, tName, data); err != nil {
		// REDIRECT to error => rendering template
		return
	}
}
