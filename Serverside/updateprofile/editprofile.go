package updateprofile

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"

	auth "Riddle_ToD/Serverside/auth"
	"Riddle_ToD/Serverside/database"
)

// ServerResponse represents the structure of the response sent to the client.
type ServerResponse struct {
	Message string `json:"message"`
}

// UserProfile holds user profile update details.
type UserProfile struct {
	NewUsername string `json:"newusername"`
	NewEmail    string `json:"newemail"`
	NewAvatar   string `json:"newavatar"`
}

// UserUpdatePassword holds user password update details.
type UserUpdatePassword struct {
	OldPassword     string `json:"oldpassword"`
	NewPassword     string `json:"newpassword"`
	ConfirmPassword string `json:"confirmpassword"`
}

// UpdateUserProfileHandler handles user profile update requests.
func UpdateUserProfileHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB, sessionStore *sessions.CookieStore) {
	// Retrieve session details.
	session, _ := sessionStore.Get(r, "session-name")
	username := session.Values["Username"]
	userID := session.Values["User_ID"]
	isAuthenticated := session.Values["Authenticated"]

	// Redirect unauthenticated users to the introduction page.
	if username == nil || isAuthenticated == false {
		http.Redirect(w, r, "/intro", http.StatusFound)
		return
	}

	// Read and parse the request body.
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var userDetails UserProfile
	if err := json.Unmarshal(body, &userDetails); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Load registered users from the database.
	auth.LoadExistingUsers(db)
	registeredUsers := auth.ExistingUsers
	isUpdateAllowed := true
	var responseMessage ServerResponse

	// Check for duplicate email or username in the system.
	for _, user := range registeredUsers {
		if user.User_ID != userID {
			if userDetails.NewEmail == user.Email {
				isUpdateAllowed = false
				responseMessage = ServerResponse{"Cannot change to this email, it already exists."}
				break
			}
			if userDetails.NewUsername == user.Username {
				isUpdateAllowed = false
				responseMessage = ServerResponse{"Cannot change to this username, it already exists."}
				break
			}
		}
	}

	// Update the user profile if no conflicts are found.
	if isUpdateAllowed {
		err = UpdateUserProfileInDatabase(db, userID, userDetails)
		if err != nil {
			http.Error(w, "Error updating user profile", http.StatusInternalServerError)
			return
		}
		responseMessage = ServerResponse{"Profile updated successfully."}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseMessage)
}

// UpdateUserProfileInDatabase updates the user profile in the database within a transaction.
func UpdateUserProfileInDatabase(db *gorm.DB, userID interface{}, userDetails UserProfile) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&database.Player{}).Where("id = ?", userID).Update("Email", userDetails.NewEmail).Error; err != nil {
			return err
		}
		if err := tx.Model(&database.Player{}).Where("id = ?", userID).Update("Username", userDetails.NewUsername).Error; err != nil {
			return err
		}
		if err := tx.Model(&database.Player{}).Where("id = ?", userID).Update("AvatarSelected", userDetails.NewAvatar).Error; err != nil {
			return err
		}
		return nil
	})
}

// UpdateUserPasswordHandler handles user password update requests.
func UpdateUserPasswordHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB, sessionStore *sessions.CookieStore) {
	// Retrieve session details.
	session, _ := sessionStore.Get(r, "session-name")
	username := session.Values["Username"]
	userID := session.Values["User_ID"]
	isAuthenticated := session.Values["Authenticated"]

	// Redirect unauthenticated users to the introduction page.
	if username == nil || isAuthenticated == false {
		http.Redirect(w, r, "/intro", http.StatusFound)
		return
	}

	// Read and parse the request body.
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var userPassword UserUpdatePassword
	if err := json.Unmarshal(body, &userPassword); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	isUpdateAllowed := true
	var responseMessage ServerResponse
	userDetails := auth.PlayerLogin{
		Username: fmt.Sprint(username),
		Password: userPassword.OldPassword,
	}

	// Validate the old password.
	if !auth.ValidateLoginCredentials(userDetails, db) {
		responseMessage = ServerResponse{"Your old password is incorrect, password update failed."}
		isUpdateAllowed = false
	}

	// Update the password if validation is successful.
	if isUpdateAllowed {
		err = UpdateUserPasswordInDatabase(db, userID, userPassword.NewPassword)
		if err != nil {
			http.Error(w, "Error updating user password", http.StatusInternalServerError)
			return
		}
		responseMessage = ServerResponse{"Password updated successfully."}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseMessage)
}

// UpdateUserPasswordInDatabase updates the user's password in the database.
func UpdateUserPasswordInDatabase(db *gorm.DB, userID interface{}, newPassword string) error {
	return db.Model(&database.Player{}).Where("id = ?", userID).Update("Password", newPassword).Error
}
