package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Player struct {
	Username string `gorm:"primaryKey"`
	Email    string `gorm:"unique"`
	Password string
}

type PlayerLevel struct {
	Level    int    `gorm:"not null"`
	Username string `gorm:"not null"` // Ensures the foreign key cannot be null
	Player   Player `gorm:"foreignKey:Username;references:Username"`
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./mydatabase"), &gorm.Config{})
	if err != nil {
		fmt.Println("error creating database")
	}
	db.AutoMigrate(&Player{})
	db.AutoMigrate(&PlayerLevel{})
	return db
}
