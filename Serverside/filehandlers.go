package Serverside

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func Filehandler(filename string, w http.ResponseWriter) {
	tmp, err := template.ParseFiles(filename)

	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		fmt.Println(err)
		os.Exit(1)
	}

	if err := tmp.Execute(w, nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		os.Exit(1)
	}

}
