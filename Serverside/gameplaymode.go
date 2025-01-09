package Serverside

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type RiddleCategoteries struct {
	Logic       bool
	Wordriddle  bool
	Mathematics bool
	General     bool
}

func Selectmode(w http.ResponseWriter, r *http.Request) {
	Filehandler("soloPlayer/GameMode.html", w)

	 vars:=mux.Vars(r)
	 path:=vars["category"]
	// if path == "" {
	// 	http.Error(w, "Bad request", http.StatusBadRequest)
	// 	return
	// }

	selectedcategory := RiddleCategoteries{}

	selectedcategory.General = false
	selectedcategory.Logic = false
	selectedcategory.Wordriddle = false
	selectedcategory.Mathematics = false

	if path == "mathematics" {
		selectedcategory.Mathematics = true
		selectedcategory.General = false
		selectedcategory.Logic = false
		selectedcategory.Wordriddle = false
	} else if path == "generalriddles" {
		selectedcategory.General = true
		selectedcategory.Logic = false
		selectedcategory.Wordriddle = false
		selectedcategory.Mathematics = false
	} else if path == "wordriddles" {
		selectedcategory.Wordriddle = true
		selectedcategory.General = false
		selectedcategory.Logic = false
		selectedcategory.Mathematics = false
	} else if path == "logicriddles" {
		selectedcategory.Logic = true
		selectedcategory.General = false
		selectedcategory.Wordriddle = false
		selectedcategory.Mathematics = false

	} else {
		http.Error(w, "NOT FOUND", http.StatusNotFound)
		return
	}
	fmt.Println(selectedcategory.Wordriddle)
	LoadRiddle(w, r, selectedcategory)
}
