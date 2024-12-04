package Serverside

import (
	"html/template"
	"net/http"
)

func Playsection(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("gamemode")

	if path == "classic" {
		temp, err := template.ParseFiles("soloPlayer/playsection.html")

		if err != nil {
			http.Error(w, "Error file", http.StatusInternalServerError)
		}

		if err = temp.Execute(w, nil); err != nil {
			http.Error(w, "Error file", http.StatusInternalServerError)
		}
	}

}
