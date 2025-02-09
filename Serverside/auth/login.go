package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	database "Riddle_ToD/Serverside/database"
)

// PlayerLogin represents the structure for login details
type PlayerLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var mu sync.Mutex

// Login handles user login requests
func Login(db *gorm.DB, w http.ResponseWriter, r *http.Request, store *sessions.CookieStore) {
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
	for i := 0; i < len(ExistingUsers); i++ {
		if ExistingUsers[i].Username == loginDetails.Username {
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

	if isUserExist && isPasswordCorrect {
		mu.Lock()
		session, _ := store.Get(r, "session-name")

		// Iterate through the session values and remove old entries for the same user
		for key, value := range session.Values {
			if value == loginDetails.Username || key == "Authenticated" {

				session.Values["Username"] = ""
				session.Values["Authenitcated"] = false
				session.Options.MaxAge = -1
				delete(session.Values, key)
			}
		}

		session.Save(r, w) // Save after deletions
		// Set new session values for the user
		session.Values["Username"] = loginDetails.Username
		session.Values["User_ID"] = FindUserId(loginDetails, db)
		session.Values["Authenticated"] = true

		// Set session options only for this user's session
		session.Options.MaxAge = 3600 // 1 hour expiration
		session.Options.HttpOnly = true
		session.Options.Secure = true
		session.Options.SameSite = http.SameSiteLaxMode

		session.Save(r, w) // Save the new session
		mu.Unlock()
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
	}

	// Compare the provided password with the stored password
	return VerifyPassword(loginDetails.Password, storedPassword)
}

func VerifyPassword(providedPassword, storedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(providedPassword))
	return err == nil
}

func FindUserId(loginDetails PlayerLogin, db *gorm.DB) string {
	var existingPlayer database.Player // Structure for database query
	var UserId string

	// Query the database for the password associated with the username
	err := db.Model(&existingPlayer).Select("User_ID").Where("Username = ?", loginDetails.Username).Scan(&UserId).Error
	if err != nil {
		log.Println("Error retrieving userId:", err)
		os.Exit(1)
	}

	return UserId
}
