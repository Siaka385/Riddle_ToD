package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"gorm.io/gorm"

	"Riddle_ToD/Serverside"
	auth "Riddle_ToD/Serverside/auth"
	database "Riddle_ToD/Serverside/database"
)

var (
	db *gorm.DB
	mu sync.Mutex
)

func InitilizeDatabase() {
	db = database.Init()
}

func Router(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path == "/" {
		Serverside.AuthenticationPageHandler(w, r, "auth_templates/loginpage.html")
	} else if path == "/gameplaymode" {
		Serverside.Selectmode(w, r)
	} else if path == "/playsection" {
		Serverside.Playsection(w, r)
	} else if path == "/DifficultySetting" {
		Serverside.Filehandler("soloPlayer/disfficult.html", w)
	} else if path == "/help" {
		Serverside.Filehandler("help/help.html", w)
	} else if path == "/register" {
		auth.Register(w, r, db)
	} else if path == "/login" {
		auth.Login(db, w, r)
	} else {
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

func main() {
	mux := http.NewServeMux()

	InitilizeDatabase()
	mux.HandleFunc("/indexFolder/", Serverside.StaticServer)
	mux.HandleFunc("/images/", Serverside.StaticServer)
	mux.HandleFunc("/js/", Serverside.StaticServer)
	mux.HandleFunc("/css/", Serverside.StaticServer)
	mux.HandleFunc("/", Router)

	fmt.Println("Server is running on port 8089...")
	if err := http.ListenAndServe(":8089", mux); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
