package Serverside

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// type IndexDisplay struct {
// 	ShowmainMenu string
// 	Showcategory string
// }

// func LoadIndexPage(w http.ResponseWriter, r *http.Request) {
// 	tmp, err := template.ParseFiles("index.html")
// 	if err != nil {
// 		http.Error(w, "Server Error", http.StatusInternalServerError)
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	k := r.URL.Query().Get("selectcategory")
// 	ted := IndexDisplay{}

// 	if k == "" {
// 		ted.ShowmainMenu = ""
// 		ted.Showcategory = "d-none"
// 	} else {
// 		ted.ShowmainMenu = "d-none"
// 		ted.Showcategory = ""
// 	}

// 	if err := tmp.Execute(w, ted); err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		os.Exit(1)
// 	}
// }

func AuthenticationPageHandler(w http.ResponseWriter, r *http.Request, temp string) {
	tmp, err := template.ParseFiles(temp)
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
