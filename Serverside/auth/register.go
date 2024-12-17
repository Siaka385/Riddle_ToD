package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"gorm.io/gorm"

	database "Riddle_ToD/Serverside/database"
)

// Renamed struct and variables
type PlayerInput struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPass"`
	Avatar          string `json:"selectedAvatar"`
}

type APIResponse struct {
	Message string `json:"response"`
}

type ExistingUser struct {
	Username string
	Email    string
}

var existingUsers []ExistingUser

func Register(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Invalid content type, expected application/json", http.StatusUnsupportedMediaType)
		return
	}

	uservalid := true

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading request body")
		os.Exit(1)
	}

	var playerInput PlayerInput
	json.Unmarshal(body, &playerInput)
	fmt.Println(playerInput)

	var apiResponse APIResponse
	apiResponse = APIResponse{"ok"}

	if playerInput.Password != playerInput.ConfirmPassword {
		apiResponse = APIResponse{"Passwords do not match"}
		uservalid = false
	}

	LoadExistingUsers(db)

	for i := 0; i < len(existingUsers); i++ {
		if existingUsers[i].Email == playerInput.Email {
			apiResponse = APIResponse{"Email already exists"}
			uservalid = false
		}
		if existingUsers[i].Username == playerInput.Username {
			apiResponse = APIResponse{"Username already exists"}
			uservalid = false
		}
	}

	if uservalid {

		AddNewUser(db, playerInput)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apiResponse)
}

func AddNewUser(db *gorm.DB, playerInput PlayerInput) error {
	var newPlayer database.Player

	newPlayer.Username = playerInput.Username
	newPlayer.Email = playerInput.Email
	newPlayer.Password = playerInput.Password

	if err := db.Create(&newPlayer).Error; err != nil {
		return err
	}

	return nil
}

func LoadExistingUsers(db *gorm.DB) {
	var playerRecord database.Player

	if err := db.Model(&playerRecord).Select("username", "email").Find(&existingUsers).Error; err != nil {
		log.Println("Error retrieving usernames and emails:", err)
	} else {
		fmt.Println("Existing users:", existingUsers)
	}
}
