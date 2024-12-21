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
	AvatarSelected string
	CreationDate   string
}

type PlayerLevel struct {
	Level   int    `gorm:"not null"`
	User_ID string `gorm:"not null"` // Ensures the foreign key cannot be null
	Player  Player `gorm:"foreignKey:Username;references:User_ID"`
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./Database/DarkRelmDatabase"), &gorm.Config{})
	if err != nil {
		fmt.Println("error creating database")
	}
	db.AutoMigrate(&Player{})
	db.AutoMigrate(&PlayerLevel{})
	return db
}
