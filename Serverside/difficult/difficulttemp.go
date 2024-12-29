package difficult

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

func DifficultHandler(w http.ResponseWriter, r *http.Request, store *sessions.CookieStore) {

	fmt.Println("Hello")
	session, _ := store.Get(r, "session-name")

	username := session.Values["Username"]
	isLogged := session.Values["Authenticated"]

	if username == nil || isLogged == false {
		http.Redirect(w, r, "/intro", http.StatusFound)
		return
	}

	tmp, err := template.ParseFiles("soloPlayer/difficult.html")
	if err != nil {
		fmt.Println("Error parsing the file:", err)
		os.Exit(1)
	}
	if err = tmp.Execute(w, nil); err != nil {
		fmt.Println("Eror executing the file")
		os.Exit(1)
	}
}
