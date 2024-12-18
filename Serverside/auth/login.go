package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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

	if isUserExist && isPasswordCorrect {
		session, _ := store.Get(r, "session-name")
		// Set some session values.
		session.Values["Username"] = loginDetails.Username
		session.Values["Authenitcated"] = true
		store.Options = &sessions.Options{
			Path:     "/",                  // The root path for the cookie (accessible from all paths)
			MaxAge:   3600,                 // Session expiration in seconds (1 hour in this case)
			HttpOnly: true,                 // Prevents JavaScript access to the cookie
			Secure:   true,                 // Ensures the cookie is sent only over HTTPS
			SameSite: http.SameSiteLaxMode, // Prevents the cookie from being sent with cross-site requests
		}

		session.Save(r, w)
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
