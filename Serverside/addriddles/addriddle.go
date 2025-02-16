package addriddles

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"Riddle_ToD/Serverside/database"
)

func Addriddle(targetDB *gorm.DB) error {
	sourceDB, err := gorm.Open(sqlite.Open("./copydb/DarkRelmDatabase"), &gorm.Config{})
	if err != nil {
		fmt.Println("error creating database")
	}
	var riddles []database.Riddle
	var choices []database.Choice
	var hints []database.Hint

	// Fetch all riddles, choices, and hints from the source database
	if err := sourceDB.Find(&riddles).Error; err != nil {
		return fmt.Errorf("error fetching riddles: %w", err)
	}
	if err := sourceDB.Find(&choices).Error; err != nil {
		return fmt.Errorf("error fetching choices: %w", err)
	}
	if err := sourceDB.Find(&hints).Error; err != nil {
		return fmt.Errorf("error fetching hints: %w", err)
	}

	// Insert data into the target database within a transaction
	return targetDB.Transaction(func(tx *gorm.DB) error {
		if len(riddles) > 0 {
			if err := tx.CreateInBatches(&riddles, 100).Error; err != nil {
				return fmt.Errorf("error adding riddles: %w", err)
			}
		}
		if len(choices) > 0 {
			if err := tx.CreateInBatches(&choices, 100).Error; err != nil {
				return fmt.Errorf("error adding choices: %w", err)
			}
		}
		if len(hints) > 0 {
			if err := tx.CreateInBatches(&hints, 100).Error; err != nil {
				return fmt.Errorf("error adding hints: %w", err)
			}
		}
		return nil // Commit transaction
	})
}
