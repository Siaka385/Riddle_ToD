package Serverside

import (
	"net/http"

	"github.com/gorilla/sessions"
)

func Logout(w http.ResponseWriter, r *http.Request, store *sessions.CookieStore) {
	session, _ := store.Get(r, "session-name")
	session.Values["Authenitcated"] = false
	session.Options.MaxAge = -1
	session.Options.HttpOnly = true // Optional: Prevent cookie access via JavaScript
	session.Options.Secure = true   // Ensure cookie is only sent over HTTPS
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)

}
