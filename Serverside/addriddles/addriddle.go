package addriddles

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"Riddle_ToD/Serverside/database"
)

func Addriddle(db *gorm.DB) {

	riddles := []database.Riddle{
		{
			ID:          1,
			Question:    "What has to be broken before you can use it?",
			Answer:      "An egg",
			Explanation: "An egg is something that must be broken before it can be used for cooking.",
			Category:    "Logic",
			Difficulty:  "Easy",
			Points:      5,
			Choices: []database.Choice{
				{ID: 1, RiddleID: 1, Text: "A coconut"},
				{ID: 2, RiddleID: 1, Text: "An egg"},
				{ID: 3, RiddleID: 1, Text: "A glass bottle"},
				{ID: 4, RiddleID: 1, Text: "A lock"},
			},
			Hints: []database.Hint{
				{ID: 1, RiddleID: 1, Text: "It’s often part of breakfast."},
				{ID: 2, RiddleID: 1, Text: "It has a shell."},
			},
		},
	}
	fmt.Println(len(riddles))
	for i := 0; i < len(riddles); i++ {
		err := db.Create(&riddles[i]).Error
		if err != nil {
			// Log the error and stop processing further
			fmt.Printf("Failed to create riddle at index %d: %v\n", i, err)
			break
		}
		time.Sleep(time.Second * 5)
	}

}
