package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"

	"Riddle_ToD/Serverside"
	auth "Riddle_ToD/Serverside/auth"
	database "Riddle_ToD/Serverside/database"
	difficulty "Riddle_ToD/Serverside/difficult"
	prof "Riddle_ToD/Serverside/updateprofile"
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

func main() {
	//mux := http.NewServeMux()

	InitilizeDatabase()

	r := mux.NewRouter()

	// Authentication and Registration Routes
	r.HandleFunc("/loginpage", Serverside.RenderAuthPage).Methods("GET")
	r.HandleFunc("/registerpage", Serverside.RenderAuthPage).Methods("GET")
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) { auth.Login(db, w, r, Store) }).Methods("POST")
	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) { auth.Register(w, r, db) }).Methods("POST")
	r.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) { auth.Logout(w, r, Store) }).Methods("POST")

	// Profile Management Routes
	r.HandleFunc("/editprofile", func(w http.ResponseWriter, r *http.Request) { prof.ProfileHandler(w, r, Store, db) }).Methods("GET")
	r.HandleFunc("/updateprofile", func(w http.ResponseWriter, r *http.Request) { prof.UpdateUserProfileHandler(w, r, db, Store) }).Methods("POST")
	r.HandleFunc("/updatepassword", func(w http.ResponseWriter, r *http.Request) { prof.UpdateUserPasswordHandler(w, r, db, Store) }).Methods("POST")

	// Static Pages
	r.HandleFunc("/intro", func(w http.ResponseWriter, r *http.Request) {
		Serverside.Filehandler("intro_template/intropage.html", w)
	}).Methods("GET")
	r.HandleFunc("/help", func(w http.ResponseWriter, r *http.Request) { Serverside.Filehandler("help/help.html", w) }).Methods("GET")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { Serverside.CheckUserSession(w, r, Store, db) }).Methods("GET")

	// Gameplay Routes
	r.HandleFunc("/gameplaymode", Serverside.Selectmode).Methods("GET")
	r.HandleFunc("/playsection", Serverside.Playsection).Methods("GET")
	//r.HandleFunc("/difficultySetting", func(w http.ResponseWriter, r *http.Request) { Serverside.Filehandler("soloPlayer/difficult.html", w) }).Methods("GET")
	r.HandleFunc("/difficultysetting", func(w http.ResponseWriter, r *http.Request) {
		difficulty.DifficultHandler(w, r, Store, db)
	}).Methods("GET")
	r.HandleFunc("/setdifficult", func(w http.ResponseWriter, r *http.Request) { difficulty.SetPlayerDifficulty(w, r, db, Store) }).Methods("POST")

	// Serve Static Files
	// Serve Static Files
	r.PathPrefix("/indexFolder/").Handler(http.StripPrefix("/indexFolder/", http.FileServer(http.Dir("indexFolder"))))
	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	//r.HandleFunc("/", Router)

	fmt.Println("Server is running on port 8089...")
	if err := http.ListenAndServe(":8089", r); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
