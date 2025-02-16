package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	User_ID  string
}

var ExistingUsers []ExistingUser

func isValidEmail(email string) (bool, error) {
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re, err := regexp.Compile(emailPattern)
	if err != nil {
		return false, err // Return error instead of crashing
	}
	return re.MatchString(email), nil
}

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

	var apiResponse APIResponse
	apiResponse = APIResponse{"ok"}

	if playerInput.Password != playerInput.ConfirmPassword {
		apiResponse = APIResponse{"Passwords do not match"}
		uservalid = false
	}

	LoadExistingUsers(db)

	for i := 0; i < len(ExistingUsers); i++ {
		if ExistingUsers[i].Email == playerInput.Email {
			apiResponse = APIResponse{"Email already exists"}
			uservalid = false
		}
		if ExistingUsers[i].Username == playerInput.Username {
			apiResponse = APIResponse{"Username already exists"}
			uservalid = false
		}
	}

	isEmailValid, err := isValidEmail(playerInput.Email)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if !isEmailValid {
		apiResponse = APIResponse{"Email is not valid"}
		uservalid = false
	}
	if uservalid {
		AddNewUser(db, playerInput)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apiResponse)
}

func AddNewUser(db *gorm.DB, playerInput PlayerInput) error {
	// Validate input
	if playerInput.Username == "" || playerInput.Email == "" || playerInput.Password == "" {
		return fmt.Errorf("invalid input: username, email, and password are required")
	}
	CurrentTime := time.Now()

	// Initialize Player and PlayerLevel
	newPlayer := database.Player{
		User_ID:        GiveUserId(),
		Username:       playerInput.Username,
		Email:          playerInput.Email,
		Password:       HashPassword(playerInput.Password), // Hash the password securely
		AvatarSelected: playerInput.Avatar,
		CreationDate:   CurrentTime.Format("2006-01-02 15:04:05"),
	}

	playlevel := database.PlayerLevel{
		Level:              1,
		PreferedDifficulty: "Easy",
		AnsweredRiddles:    []uint{},
		User_ID:            newPlayer.User_ID,
	}

	playerEasyStat := database.EasyStat{
		User_ID:    newPlayer.User_ID,
		General:    0,
		Logic:      0,
		Mathematic: 0,
		Word:       0,
		Total:      0,
	}

	playerMediumStat := database.MediumStat{
		User_ID:    newPlayer.User_ID,
		General:    0,
		Logic:      0,
		Mathematic: 0,
		Word:       0,
		Total:      0,
	}

	playerHardStat := database.HardStat{
		User_ID:    newPlayer.User_ID,
		General:    0,
		Logic:      0,
		Mathematic: 0,
		Word:       0,
		Total:      0,
	}

	// Use a transaction for atomicity
	//A transaction in a database is a sequence of operations (like INSERT, UPDATE, or DELETE) that are executed as a single logical unit.
	// If all the operations within the transaction succeed, the changes are permanently saved to the database (committed).
	// If any operation fails, all changes made during the transaction are undone (rolled back).
	return db.Transaction(func(tx *gorm.DB) error {
		// Create the new player
		if err := tx.Create(&newPlayer).Error; err != nil {
			return fmt.Errorf("failed to create player: %w", err)
		}

		// Initialize the player's level
		if err := tx.Create(&playlevel).Error; err != nil {
			return fmt.Errorf("failed to create player level: %w", err)
		}

		if err := tx.Create(&playerEasyStat).Error; err != nil {
			return fmt.Errorf("failed to create player easy stat: %w", err)
		}

		if err := tx.Create(&playerMediumStat).Error; err != nil {
			return fmt.Errorf("failed to create player medium stat: %w", err)
		}

		if err := tx.Create(&playerHardStat).Error; err != nil {
			return fmt.Errorf("failed to create Hard stat: %w", err)
		}

		return nil
	})
}

func LoadExistingUsers(db *gorm.DB) {
	var playerRecord database.Player

	if err := db.Model(&playerRecord).Select("username", "email", "User_ID").Find(&ExistingUsers).Error; err != nil {
		log.Println("Error retrieving usernames,emails and User_ID:", err)
	}
}

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password")
		return ""
	}
	return string(hashedPassword)
}

func GiveUserId() string {
	userid := uuid.New()
	return userid.String()
}
