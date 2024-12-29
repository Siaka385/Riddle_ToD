package difficult

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"

	"Riddle_ToD/Serverside/database"
)

type Difficulty struct {
	SelectedDifficulty string `json:"selectedDifficulty"`
}

func SetPlayerDifficulty(w http.ResponseWriter, r *http.Request, db *gorm.DB, store *sessions.CookieStore) {
	session, _ := store.Get(r, "session-name")

	username := session.Values["Username"]
	isAuthenticated := session.Values["Authenticated"]
	userID := session.Values["User_ID"]

	if username == nil || isAuthenticated == false {
		http.Redirect(w, r, "/intro", http.StatusFound)
		return
	}

	var difficulty Difficulty

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading the request body:", err)
		http.Error(w, "Unable to process the request", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &difficulty); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		os.Exit(1)
		return
	}

	err = UpdatePreferredDifficulty(userID, db, difficulty.SelectedDifficulty)
	if err != nil {
		fmt.Println("Error updating preferred difficulty:", err)
		http.Error(w, "Unable to set difficulty", http.StatusInternalServerError)
		os.Exit(1)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func UpdatePreferredDifficulty(userID any, db *gorm.DB, difficulty string) error {
	return db.Model(&database.PlayerLevel{}).
		Where("User_ID = ?", userID).
		Update("PreferedDifficulty", difficulty).
		Error
}

func FetchPreferredDifficulty(db *gorm.DB, userID any) (string, error) {
	var preferredDifficulty string
	err := db.Model(&database.PlayerLevel{}).
		Select("PreferedDifficulty").
		Where("User_ID = ?", userID).
		Scan(&preferredDifficulty).
		Error

	if err != nil {
		fmt.Println("Error fetching the preferred difficulty level:", err)
		return "", err
	}

	return preferredDifficulty, nil
}
