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
		Filehandler("index.html", w)
	} else if path == "/gameplaymode" {
		Selectmode(w, r)
	} else {
		http.Error(w, "Not found", http.StatusNotFound)
	}
}
