package Serverside

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
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

func CheckUserSession(w http.ResponseWriter, r *http.Request, Store *sessions.CookieStore, db *gorm.DB) {
	session, _ := Store.Get(r, "session-name")

	username := session.Values["Username"]
	isLogged := session.Values["Authenticated"]

	if username == nil || isLogged == false {
		http.Redirect(w, r, "/intro", http.StatusFound)
		return
	}
	http.Redirect(w, r, "/Dashboard", http.StatusFound)
}

func RenderAuthPage(responseWriter http.ResponseWriter, request *http.Request) {
	var templatePath string
	if request.URL.Path == "/loginpage" {
		templatePath = "auth_templates/loginpage.html"
	} else {
		templatePath = "auth_templates/signuppage.html"
	}

	templateFile, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(responseWriter, "Server Error", http.StatusInternalServerError)
		fmt.Println(err)
		os.Exit(1)
	}

	if err := templateFile.Execute(responseWriter, nil); err != nil {
		http.Error(responseWriter, "Internal Server Error", http.StatusInternalServerError)
		os.Exit(1)
	}
}
