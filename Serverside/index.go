package Serverside

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"

	database "Riddle_ToD/Serverside/database"
)

type IndexContent struct {
	UserName      string
	UserLevel     int
	LevelAlias    string
	MainMenuClass string
	CategoryClass string
	AvatarIcon    string
}

type UserLevelInfo struct {
	CurrentLevel int
	LevelAlias   string
}

func RenderIndexPage(responseWriter http.ResponseWriter, request *http.Request, db *gorm.DB, sessionn *sessions.CookieStore) {
	templateFile, err := template.ParseFiles("UserDashboard.html")
	if err != nil {
		//http.Error(responseWriter, "Server Error", http.StatusInternalServerError)
		fmt.Println("template error", err)
		os.Exit(1)
	}

	session, _ := sessionn.Get(request, "session-name")

	//username := session.Values["Username"]
	userId := session.Values["User_ID"]

	category := request.URL.Query().Get("selectcategory")
	pageContent := IndexContent{}
	userlevel, err := FetchUserLevel(db, userId)
	if err != nil {
		//	http.Error(responseWriter, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("fetch error", err)
		os.Exit(1)
	}

	pageContent.UserName, err = FetchUsername(db, userId)
	if err != nil {
		//	http.Error(responseWriter, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("fetch username:", err)
		os.Exit(1)
	}
	pageContent.UserLevel = userlevel.CurrentLevel
	pageContent.LevelAlias = userlevel.LevelAlias
	pageContent.AvatarIcon, err = FetchAvatarIcon(db, userId)
	if err != nil {
		//	http.Error(responseWriter, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("fetch icon", err)
		os.Exit(1)
	}

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

func FetchUserLevel(db *gorm.DB, userids any) (UserLevelInfo, error) {
	var userLevel UserLevelInfo
	fmt.Println(userids)
	// Perform the query
	if err := db.Model(&database.PlayerLevel{}).
		Select("Level").
		Where("User_ID = ?", userids).
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

func FetchAvatarIcon(db *gorm.DB, userids any) (string, error) {
	var avatarIcon string
	fmt.Println(userids)
	// Perform the query
	if err := db.Model(&database.Player{}).
		Select("AvatarSelected").
		Where("User_ID   = ?", userids).
		Scan(&avatarIcon).Error; err != nil {
		return avatarIcon, fmt.Errorf("failed to fetch user level: %w", err)
	}

	fmt.Println("AVATAR:", avatarIcon)
	return avatarIcon, nil
}

func FetchUsername(db *gorm.DB, userids any) (string, error) {
	var Username string

	// Perform the query
	if err := db.Model(&database.Player{}).
		Select("Username").
		Where("User_ID   = ?", userids).
		Scan(&Username).Error; err != nil {
		return Username, fmt.Errorf("failed to fetch user level: %w", err)
	}
	fmt.Println("username:", Username)
	return Username, nil

}
