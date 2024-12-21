package updateprofile

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"

	"Riddle_ToD/Serverside/auth"
)

type UserProfiles struct {
	Username  string
	Useremail string
}

func ProfileHandler(w http.ResponseWriter, r *http.Request, store *sessions.CookieStore, db *gorm.DB) {

	session, _ := store.Get(r, "session-name")
	username := session.Values["Username"]
	userID := session.Values["User_ID"]
	isAuthenticated := session.Values["Authenticated"]

	// Redirect unauthenticated users to the introduction page.
	if username == nil || isAuthenticated == false {
		http.Redirect(w, r, "/intro", http.StatusFound)
		return
	}
	var userprofile UserProfiles
	auth.LoadExistingUsers(db)
	users := auth.ExistingUsers
	for i := 0; i < len(users); i++ {
		if users[i].User_ID == userID {
			userprofile.Username = users[i].Username
			userprofile.Useremail = users[i].Email
		}
	}
	tmp, err := template.ParseFiles("UserProfileEdit_template/UserProfile.html")
	if err != nil {
		fmt.Println("ERROR PARSING FILE")
		os.Exit(1)
	}

	if err = tmp.Execute(w, userprofile); err != nil {
		fmt.Println("ERROR executing FILE")
		os.Exit(1)
	}
}
