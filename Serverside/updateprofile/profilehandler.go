package updateprofile

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"

	"Riddle_ToD/Serverside"
	"Riddle_ToD/Serverside/auth"
	"Riddle_ToD/Serverside/database"
)

type UserProfiles struct {
	UserAvatar string
	Username   string
	Useremail  string
	EasyStat   Stat
	MediumStat Stat
	HardStat   Stat
}

type Stat struct {
	General    uint
	Logic      uint
	Mathematic uint
	Word       uint
	Total      uint
}

func ProfileHandler(w http.ResponseWriter, r *http.Request, store *sessions.CookieStore, db *gorm.DB) {

	session, _ := store.Get(r, "session-name")
	username := session.Values["Username"]
	userID := session.Values["User_ID"]
	isAuthenticated := session.Values["Authenticated"]

	// Redirect unauthenticated users to the introduction page.
	if username == nil || isAuthenticated == false {
		http.Redirect(w, r, "/intro", http.StatusFound)
		return
	}
	var userprofile UserProfiles
	auth.LoadExistingUsers(db)
	users := auth.ExistingUsers

	for i := 0; i < len(users); i++ {
		if users[i].User_ID == userID {
			userprofile.Username = users[i].Username
			userprofile.Useremail = users[i].Email
			userprofile.UserAvatar, _ = Serverside.FetchAvatarIcon(db, userID)
			Stat, err := GetEasyStat(db, userID)
			if err != nil {
				fmt.Println("Failed to fetch easy stats:", err)
			} else {
				userprofile.EasyStat = Stat
			}
			Stat,err=GetMediumStat(db,userID)
			if err != nil {
				fmt.Println("Failed to fetch easy stats:", err)
			} else {
				userprofile.MediumStat = Stat
			}
			Stat,err=GetHardStat(db,userID)
			if err != nil {
				fmt.Println("Failed to fetch easy stats:", err)
			} else {
				userprofile.HardStat=Stat

		}
	}
}

	tmp, err := template.ParseFiles("UserProfileEdit_template/UserProfile.html")
	if err != nil {
		fmt.Println("ERROR PARSING FILE")
		os.Exit(1)
	}

	if err = tmp.Execute(w, userprofile); err != nil {
		fmt.Println("ERROR executing FILE")
		os.Exit(1)
	}
}


func GetEasyStat(db *gorm.DB, userID any) (Stat, error) {
	var stats Stat

	if err := db.Model(&database.EasyStat{}).
		Select("General, Logic, Mathematic, Word, Total").
		Where("User_ID = ?", userID).
		Scan(&stats).Error; err != nil {
		return stats, fmt.Errorf("failed to fetch easy stats: %w", err)
	}

	fmt.Println("Fetched Easy Stats:", stats)
	return stats, nil
}

func GetMediumStat(db *gorm.DB, userID any) (Stat, error) {
	var stats Stat

	if err := db.Model(&database.MediumStat{}).
		Select("General, Logic, Mathematic, Word, Total").
		Where("User_ID = ?", userID).
		Scan(&stats).Error; err != nil {
		return stats, fmt.Errorf("failed to fetch medium stats: %w", err)
	}

	fmt.Println("Fetched Medium Stats:", stats)
	return stats, nil
}

func GetHardStat(db *gorm.DB, userID any) (Stat, error) {
	var stats Stat

	if err := db.Model(&database.HardStat{}).
		Select("General, Logic, Mathematic, Word, Total").
		Where("User_ID = ?", userID).
		Scan(&stats).Error; err != nil {
		return stats, fmt.Errorf("failed to fetch hard stats: %w", err)
	}

	fmt.Println("Fetched Hard Stats:", stats)
	return stats, nil
}
