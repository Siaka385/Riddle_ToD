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
			ID:          195,
			Question:    "What is one thing that all people, regardless of their politics or religion, have to agree is between heaven and earth?",
			Answer:      "The word 'and.'",
			Explanation: "The word 'and' is literally between 'heaven' and 'earth.'",
			Category:    "Wordplay",
			Difficulty:  "Medium",
			Points:      10,
			Choices: []database.Choice{
				{ID: 662, RiddleID: 195, Text: "The word 'and'"},
				{ID: 663, RiddleID: 195, Text: "Air"},
				{ID: 664, RiddleID: 195, Text: "Clouds"},
				{ID: 665, RiddleID: 195, Text: "Atmosphere"},
			},
			Hints: []database.Hint{
				{ID: 232, RiddleID: 195, Text: "It's a linguistic answer."},
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
