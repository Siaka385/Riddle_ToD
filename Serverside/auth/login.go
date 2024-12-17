package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"gorm.io/gorm"
)

type PlayerLogin struct {
	Username string `json:"username"`

	Password string `json:"password"`
}

func Login(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Invalid content type, expected application/json", http.StatusUnsupportedMediaType)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading request body")
		os.Exit(1)
	}
	var playerloginDetail PlayerLogin
	json.Unmarshal(body, &playerloginDetail)

	LoadExistingUsers(db)

	for i := 0; i < len(existingUsers); i++ {

		if existingUsers[i].Username == playerInput.Username {
			apiResponse = APIResponse{"Username already exists"}
			uservalid = false
		}
	}
}
