package riddleloader

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

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
	userID := session.Values["User_ID"]
	username := session.Values["Username"]
	isAuthenticated := session.Values["Authenticated"]

	// Redirect to intro page if user is not logged in
	if username == nil || isAuthenticated == false {
		http.Redirect(w, r, "/intro", http.StatusFound)
	}

	// Get riddles the user has already solved
	solvedRiddles := GetSolvedRiddles(db, userID)

	// Get the user's preferred difficulty level
	userPreferredDifficulty := GetUserPreferredDifficulty(db, userID)
	riddlesToDisplay = []ServerResponse{}

	if len(FilteredRiddles) == 0 {
		fmt.Println("CHE")
		os.Exit(1)
	}
	fmt.Println("hello")
	for i := 0; i < len(FilteredRiddles); i++ {
		// Check if the riddle is unsolved and matches the user's preferred difficulty
		if !IsRiddleSolved(solvedRiddles, FilteredRiddles[i].ID) &&
			strings.EqualFold(FilteredRiddles[i].Difficulty, userPreferredDifficulty) {
			k := ServerResponse{
				Id:      uint(FilteredRiddles[i].ID),
				Riddle:  FilteredRiddles[i].Question,
				Choices: StringIT(RiddleChoicesMap[FilteredRiddles[i].ID]),
				Hint:    ReleaseHint(RiddleHintsMap[FilteredRiddles[i].ID]),
			}
			riddlesToDisplay = append(riddlesToDisplay, k)
		}

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

func GetSolvedRiddles(db *gorm.DB, userID any) []string {
	var solvedRiddlesString string

	// Fetch solved riddles from the database
	err := db.Model(&database.PlayerLevel{}).
		Select("AnsweredRiddles").
		Where("user_id = ?", userID).
		Scan(&solvedRiddlesString).Error
	if err != nil {
		fmt.Println("Error fetching solved riddles:", err)
	}

	// Split the solved riddles string into a slice
	return strings.Split(solvedRiddlesString, " ")
}

func GetUserPreferredDifficulty(db *gorm.DB, userID any) string {
	var preferredDifficulty string

	// Fetch the user's preferred difficulty from the database
	err := db.Model(&database.PlayerLevel{}).
		Select("prefered_difficulty").
		Where("user_id = ?", userID).
		Scan(&preferredDifficulty).Error
	if err != nil {
		fmt.Println("Error fetching preferred difficulty:", err)
	}

	return preferredDifficulty
}

func IsRiddleSolved(solvedRiddles []string, riddleID int) bool {
	for _, solvedRiddle := range solvedRiddles {
		id, _ := strconv.Atoi(solvedRiddle)

		if riddleID == id {
			return true
		}
	}
	return false
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
	fmt.Println("request", ted.Count)
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
