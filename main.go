package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"

	"Riddle_ToD/Serverside"
	auth "Riddle_ToD/Serverside/auth"
	database "Riddle_ToD/Serverside/database"
	utils "Riddle_ToD/Serverside/utils"
)

var (
	db *gorm.DB
	mu sync.Mutex
)

// Note: Don't store your key in your source code. Pass it via an
// environmental variable, or flag (or both), and don't accidentally commit it
// alongside your code. Ensure your key is sufficiently random - i.e. use Go's
// crypto/rand or securecookie.GenerateRandomKey(32) and persist the result.
var Store *sessions.CookieStore

func InitilizeDatabase() {
	db = database.Init()
	Store = sessions.NewCookieStore([]byte(utils.GenerateRandomString(40)))
}

func Router(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path == "/loginpage" || path == "/registerpage" {
		Serverside.RenderAuthPage(w, r)
	} else if path == "/intro" {
		Serverside.Filehandler("intro_template/intropage.html", w)
	} else if path == "/" {
		Serverside.CheckUserSession(w, r, Store, db)
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
		auth.Login(db, w, r, Store)
	} else if path == "/logout" {
		auth.Logout(w, r, Store)
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
