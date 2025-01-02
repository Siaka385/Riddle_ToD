package difficult

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

type PlayerDifficultys struct{
	Difficult string
}

func DifficultHandler(w http.ResponseWriter, r *http.Request, store *sessions.CookieStore,db *gorm.DB) {

	fmt.Println("Hello")
	session, _ := store.Get(r, "session-name")

	username := session.Values["Username"]
	isLogged := session.Values["Authenticated"]
	userID := session.Values["User_ID"]

	if username == nil || isLogged == false {
		http.Redirect(w, r, "/intro", http.StatusFound)
		return
	}
	var diff PlayerDifficultys
	diff.Difficult,_=FetchPreferredDifficulty(db,userID)
        
	tmp, err := template.ParseFiles("soloPlayer/difficult.html")
	if err != nil {
		fmt.Println("Error parsing the file:", err)
		os.Exit(1)
	}
	if err = tmp.Execute(w, diff); err != nil {
		fmt.Println("Eror executing the file")
		os.Exit(1)
	}
}
