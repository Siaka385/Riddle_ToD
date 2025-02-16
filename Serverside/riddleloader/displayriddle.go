package riddleloader

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"

	"Riddle_ToD/Serverside/database"
)

type ServerResponse struct {
	Id      uint
	Riddle  string
	Choices []string
	Hint    string
}

type RiddleNum struct {
	Count int `json:"count"`
}

var riddlesToDisplay []ServerResponse

func LoadRiddles(w http.ResponseWriter, r *http.Request, db *gorm.DB, store *sessions.CookieStore) {

	session, _ := store.Get(r, "session-name")

	// Retrieve user session details
	//userID := session.Values["User_ID"]
	username := session.Values["Username"]
	isAuthenticated := session.Values["Authenticated"]

	// Redirect to intro page if user is not logged in
	if username == nil || isAuthenticated == false {
		http.Redirect(w, r, "/intro", http.StatusFound)
	}

	// Get the user's preferred difficulty level
	riddlesToDisplay = []ServerResponse{}

	if len(FilteredRiddles) == 0 {
		os.Exit(1)
	}
	for i := 0; i < len(FilteredRiddles); i++ {
		k := ServerResponse{
			Id:      uint(FilteredRiddles[i].ID),
			Riddle:  FilteredRiddles[i].Question,
			Choices: StringIT(RiddleChoicesMap[FilteredRiddles[i].ID]),
			Hint:    ReleaseHint(RiddleHintsMap[FilteredRiddles[i].ID]),
		}
		riddlesToDisplay = append(riddlesToDisplay, k)

	}

	temp, err := template.ParseFiles("soloPlayer/playsection.html")
	if err != nil {
		http.Error(w, "Error file", http.StatusInternalServerError)
	}

	if err = temp.Execute(w, nil); err != nil {
		http.Error(w, "Error file", http.StatusInternalServerError)
	}

	// Implement rendering or response handling for `riddlesToDisplay` here
}

func GetSolvedRiddles(db *gorm.DB, userID any) []uint {
	var solvedRiddlesString []uint

	// Fetch solved riddles from the database
	err := db.Model(&database.PlayerLevel{}).
		Select("AnsweredRiddles").
		Where("user_id = ?", userID).
		Scan(&solvedRiddlesString).Error
	if err != nil {
		fmt.Println("Error fetching solved riddles:", err)
	}

	return solvedRiddlesString
}

func RiddleLoaders(w http.ResponseWriter, r *http.Request, store *sessions.CookieStore) {
	session, _ := store.Get(r, "session-name")

	// Retrieve user session details
	//userID := session.Values["User_ID"]
	username := session.Values["Username"]
	isAuthenticated := session.Values["Authenticated"]

	// Redirect to intro page if user is not logged in
	if username == nil || isAuthenticated == false {
		http.Redirect(w, r, "/intro", http.StatusFound)
		return
	}
	// Ensure we close the request body
	defer r.Body.Close()

	var ted RiddleNum
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	json.Unmarshal(body, &ted)
	if len(riddlesToDisplay) > 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(riddlesToDisplay[ted.Count])
	} else {
		fmt.Println(riddlesToDisplay)
	}
}

func StringIT(m []Choice) []string {
	var k []string
	for i := 0; i < len(m); i++ {
		k = append(k, m[i].Text)
	}

	return k
}

func ReleaseHint(m []Hint) string {
	if len(m) < 1 {
		return "No hint available"
	}
	return m[0].Text
}
