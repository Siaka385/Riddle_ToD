package Serverside

import (
	"fmt"
	"net/http"
)

type RiddleCategoteries struct {
	Logic       bool
	Wordriddle  bool
	Mathematics bool
	General     bool
}

func Selectmode(w http.ResponseWriter, r *http.Request) {
	Filehandler("soloPlayer/GameMode.html", w)

	path := r.URL.Query().Get("category")
	if path == "" {
		fmt.Println("hhyh")
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	selectedcategory := &RiddleCategoteries{}

	if path == "mathematics" {
		selectedcategory.Mathematics = true
	} else if path == "generalriddles" {
		selectedcategory.General = true
	} else if path == "wordriddles" {
		selectedcategory.Wordriddle = true
	} else if path == "logicriddles" {
		selectedcategory.Wordriddle = true
	} else {
		http.Error(w, "NOT FOUND", http.StatusNotFound)
		return
	}

}
