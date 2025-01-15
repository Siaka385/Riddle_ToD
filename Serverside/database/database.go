package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Player struct {
	User_ID        string `gorm:"primaryKey"`
	Username       string
	Email          string `gorm:"unique"`
	Password       string
	AvatarSelected string `gorm:"not null"`
	CreationDate   string
}
type PlayerLevel struct {
	Level              int `gorm:"not null"`
	PreferedDifficulty string
	AnsweredRiddles    string
	User_ID            string `gorm:"not null"`                              // Foreign key referring to Player.User_ID
	Player             Player `gorm:"foreignKey:User_ID;references:User_ID"` // Foreign key constraint on User_ID
}

type Riddle struct {
	ID          uint     `gorm:"primaryKey"`
	Question    string   `gorm:"not null"`
	Answer      string   `gorm:"not null"`
	Explanation string   `gorm:"not null"`
	Category    string   `gorm:"not null"`            // Category like "Logic", "Mathematics"
	Difficulty  string   `gorm:"not null"`            // Difficulty level like "Easy"
	Points      int      `gorm:"default:0"`           // Points for solving
	Choices     []Choice `gorm:"foreignKey:RiddleID"` // Relation to Choices
	Hints       []Hint   `gorm:"foreignKey:RiddleID"` // Relation to Hints
}
type Choice struct {
	ID       uint   `gorm:"primaryKey"`
	RiddleID uint   `gorm:"not null"` // Foreign key to Riddle
	Text     string `gorm:"not null"` // Choice text
}

type Hint struct {
	ID       uint   `gorm:"primaryKey"`
	RiddleID uint   `gorm:"not null"` // Foreign key to Riddle
	Text     string `gorm:"not null"` // Hint text
}

type EasyStat struct {
	User_ID    string `gorm:"not null"`
	General    uint   `gorm:"default:0"`
	Logic      uint   `gorm:"default:0"`
	Mathematic uint   `gorm:"default:0"`
	Word       uint   `gorm:"default:0"`
	Total      uint   `gorm:"default:0"`
	Player     Player `gorm:"foreignKey:User_ID;references:User_ID"` // Foreign key constraint on User_ID
}
type MediumStat struct {
	User_ID    string `gorm:"not null"`
	General    uint   `gorm:"default:0"`
	Logic      uint   `gorm:"default:0"`
	Mathematic uint   `gorm:"default:0"`
	Word       uint   `gorm:"default:0"`
	Total      uint   `gorm:"default:0"`
	Player     Player `gorm:"foreignKey:User_ID;references:User_ID"` // Foreign key constraint on User_ID
}
type HardStat struct {
	User_ID    string `gorm:"not null"`
	General    uint   `gorm:"default:0"`
	Logic      uint   `gorm:"default:0"`
	Mathematic uint   `gorm:"default:0"`
	Word       uint   `gorm:"default:0"`
	Total      uint   `gorm:"default:0"`
	Player     Player `gorm:"foreignKey:User_ID;references:User_ID"` // Foreign key constraint on User_ID
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./Database/DarkRelmDatabase"), &gorm.Config{})
	if err != nil {
		fmt.Println("error creating database")
	}
	db.AutoMigrate(&Player{})
	db.AutoMigrate(&PlayerLevel{})
	db.AutoMigrate(&Riddle{})
	db.AutoMigrate(&Choice{})
	db.AutoMigrate(&Hint{})
	db.AutoMigrate(&EasyStat{})
	db.AutoMigrate(&MediumStat{})
	db.AutoMigrate(&HardStat{})
	return db
}
