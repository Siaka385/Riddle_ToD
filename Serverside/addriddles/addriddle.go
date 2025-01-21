package addriddles

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"

	"Riddle_ToD/Serverside/database"
)

func Addriddle(db *gorm.DB) {

	riddles := []database.Riddle{
		{
			ID:          293,
			Question:    "The stack just might be sent all over. Full of what's new, yet it's nearly obsolete. What am I?",
			Answer:      "Newspapers",
			Explanation: "Newspapers, once a primary source for news, are now less common due to digital media.",
			Category:    "Objects",
			Difficulty:  "Medium",
			Points:      10,
			Choices: []database.Choice{
				{ID: 1054, RiddleID: 293, Text: "Newspapers"},
				{ID: 1055, RiddleID: 293, Text: "Magazines"},
				{ID: 1056, RiddleID: 293, Text: "Letters"},
				{ID: 1057, RiddleID: 293, Text: "Packages"},
			},
			Hints: []database.Hint{
				{ID: 330, RiddleID: 293, Text: "It's a source of daily information."},
			},
		},
	}
	fmt.Println(len(riddles))
	count := 0
	for i := 0; i < len(riddles); i++ {

		if riddles[i].Category != "Logic" || riddles[i].Category != "General" || !strings.Contains(riddles[i].Category, "Word") {
			count++
			if count < 30 {
				riddles[i].Category = "Logic"
			} else {
				riddles[i].Category = "General"
			}
		}
		if riddles[i].Difficulty == "Easy" {
			riddles[i].Points = 5
		}
		if riddles[i].Difficulty == "Medium" {
			riddles[i].Points = 10
		}
		if riddles[i].Difficulty == "Hard" {
			riddles[i].Points = 15
		}

		err := db.Create(&riddles[i]).Error
		if err != nil {
			// Log the error and stop processing further
			fmt.Printf("Failed to create riddle at index %d: %v\n", i, err)
			break
		}
		time.Sleep(time.Second * 3)
	}

}
