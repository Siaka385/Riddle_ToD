package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	database "Riddle_ToD/Serverside/database"
)

// PlayerLogin represents the structure for login details
type PlayerLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login handles user login requests
func Login(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// Validate that the content type is application/json
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Invalid content type, expected application/json", http.StatusUnsupportedMediaType)
		return
	}

	// Read and parse the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading request body")
		os.Exit(1)
	}
	var loginDetails PlayerLogin
	json.Unmarshal(body, &loginDetails)

	// Load all existing users into memory
	LoadExistingUsers(db)

	// Check if the username exists and validate the password
	var isPasswordCorrect bool
	isUserExist := false
	for i := 0; i < len(existingUsers); i++ {
		if existingUsers[i].Username == loginDetails.Username {
			isUserExist = true
			isPasswordCorrect = ValidateLoginCredentials(loginDetails, db)
		}
	}

	// Prepare the response
	var response APIResponse

	response = APIResponse{"ok"}
	if !isUserExist {
		response = APIResponse{"Username does not exist"}
	} else if !isPasswordCorrect {
		response = APIResponse{"Incorrect password"}
	}

	// Return the response as JSON
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// ValidateLoginCredentials checks if the username and password match
func ValidateLoginCredentials(loginDetails PlayerLogin, db *gorm.DB) bool {
	var existingPlayer database.Player // Structure for database query
	var storedPassword string

	// Query the database for the password associated with the username
	err := db.Model(&existingPlayer).Select("Password").Where("Username = ?", loginDetails.Username).Scan(&storedPassword).Error
	if err != nil {
		log.Println("Error retrieving password:", err)
		return false
	} else {
		fmt.Println("Password found:", storedPassword)
	}

	// Compare the provided password with the stored password
	return VerifyPassword(loginDetails.Password, storedPassword)
}

func VerifyPassword(providedPassword, storedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(providedPassword))
	return err == nil
}
