package riddleloader

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

type Url struct {
	Riddletype string
	Classic    string
	Timeattack string
	Survival   string
}

func Selectmode(w http.ResponseWriter, r *http.Request, db *gorm.DB, session *sessions.CookieStore) {
	sessions, _ := session.Get(r, "session-name")

	//username := session.Values["Username"]
	//	userId := sessions.Values["User_ID"]
	username := sessions.Values["Username"]
	isLogged := sessions.Values["Authenticated"]

	if username == nil || isLogged == false {
		http.Redirect(w, r, "/intro", http.StatusFound)
		return
	}

	vars := mux.Vars(r)
	path := vars["category"]

	ren := Url{
		Riddletype: strings.ToUpper(path),
		Classic:    path + "/classic",
		Timeattack: path + "/timeattack",
		Survival:   path + "/survival",
	}

	temp, err := template.ParseFiles("soloPlayer/GameMode.html")
	if err != nil {
		fmt.Println("Error parsing files")
		os.Exit(1)
	}
	temp.Execute(w, ren)

	// if path == "" {
	// 	http.Error(w, "Bad request", http.StatusBadRequest)
	// 	return
	// }

	if path == "mathematics" {
		LoadRiddlesFromDB(w, r, "Mathematics", db)
	} else if path == "generalriddles" {
		LoadRiddlesFromDB(w, r, "General", db)
	} else if path == "wordriddles" {
		LoadRiddlesFromDB(w, r, "Word", db)
	} else if path == "logicriddles" {
		LoadRiddlesFromDB(w, r, "Logic", db)
	} else {
		http.Error(w, "NOT FOUND", http.StatusNotFound)
		return
	}
}
