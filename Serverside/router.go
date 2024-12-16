package Serverside

import (
	"net/http"
	"os"
)

func StaticServer(w http.ResponseWriter, r *http.Request) {
	filePath := "." + r.URL.Path

	info, err := os.Stat(filePath)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	if info.IsDir() {
		http.Error(w, "Access forbidden", http.StatusForbidden)
		return
	}
	http.ServeFile(w, r, filePath)
}

func Router(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path == "/" {
		LoadIndexPage(w, r)
	} else if path == "/gameplaymode" {
		Selectmode(w, r)
	} else if path == "/playsection" {
		Playsection(w, r)
	} else if path == "/DifficultySetting" {
		Filehandler("soloPlayer/disfficult.html", w)
	} else if path == "/help" {
		Filehandler("help/help.html", w)
	} else {
		http.Error(w, "Not found", http.StatusNotFound)
	}
}
