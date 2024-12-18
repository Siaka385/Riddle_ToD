package Serverside

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"gorm.io/gorm"

	database "Riddle_ToD/Serverside/database"
)

type IndexContent struct {
	UserName      string
	UserLevel     int
	LevelAlias    string
	MainMenuClass string
	CategoryClass string
}

type UserLevelInfo struct {
	CurrentLevel int
	LevelAlias   string
}

func RenderIndexPage(responseWriter http.ResponseWriter, request *http.Request, db *gorm.DB, username any) {
	templateFile, err := template.ParseFiles("UserDashboard.html")
	if err != nil {
		//http.Error(responseWriter, "Server Error", http.StatusInternalServerError)
		fmt.Println("template error", err)
		os.Exit(1)
	}

	category := request.URL.Query().Get("selectcategory")
	pageContent := IndexContent{}
	userlevel, err := FetchUserLevel(db, username)
	if err != nil {
		//	http.Error(responseWriter, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("fetch error", err)
		os.Exit(1)
	}

	pageContent.UserName = fmt.Sprint(username)
	pageContent.UserLevel = userlevel.CurrentLevel
	pageContent.LevelAlias = userlevel.LevelAlias

	if category == "" {
		pageContent.MainMenuClass = ""
		pageContent.CategoryClass = "d-none"
	} else {
		pageContent.MainMenuClass = "d-none"
		pageContent.CategoryClass = ""
	}

	if err := templateFile.Execute(responseWriter, pageContent); err != nil {
		//http.Error(responseWriter, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("template execut error", err)
		os.Exit(1)
	}
}

func FetchUserLevel(db *gorm.DB, username any) (UserLevelInfo, error) {
	var userLevel UserLevelInfo

	// Perform the query
	if err := db.Model(&database.PlayerLevel{}).
		Select("Level").
		Where("username = ?", username).
		Scan(&userLevel.CurrentLevel).Error; err != nil {
		return UserLevelInfo{}, fmt.Errorf("failed to fetch user level: %w", err)
	}

	// Assign level nickname based on level range
	switch {
	case userLevel.CurrentLevel >= 1 && userLevel.CurrentLevel <= 10:
		userLevel.LevelAlias = "Novice Riddler"
	case userLevel.CurrentLevel >= 11 && userLevel.CurrentLevel <= 20:
		userLevel.LevelAlias = "Intermediate Riddler"
	case userLevel.CurrentLevel >= 21 && userLevel.CurrentLevel <= 30:
		userLevel.LevelAlias = "Advanced Riddler"
	case userLevel.CurrentLevel >= 31 && userLevel.CurrentLevel <= 40:
		userLevel.LevelAlias = "Expert Riddler"
	case userLevel.CurrentLevel >= 41 && userLevel.CurrentLevel <= 50:
		userLevel.LevelAlias = "Master Riddler"
	case userLevel.CurrentLevel > 50:
		userLevel.LevelAlias = "Legendary Riddler"
	default:
		userLevel.LevelAlias = "Unknown Riddler"
	}

	return userLevel, nil
}
