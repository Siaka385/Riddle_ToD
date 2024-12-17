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

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./mydatabase"), &gorm.Config{})
	if err != nil {
		fmt.Println("error creating database")
	}
	db.AutoMigrate(&Player{})
	return db
}
