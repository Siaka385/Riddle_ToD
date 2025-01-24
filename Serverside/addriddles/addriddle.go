package addriddles

// import (
// 	"fmt"
// 	"strings"
// 	"time"
//
//
//
//

// 	"gorm.io/gorm"

// 	"Riddle_ToD/Serverside/database"
// )

// func Addriddle(db *gorm.DB) {

// 		//riddle 303
// 	}
// 	fmt.Println(len(riddles))
// 	count := 0
// 	for i := 0; i < len(riddles); i++ {

// 		if riddles[i].Category != "Logic" || riddles[i].Category != "General" || !strings.Contains(riddles[i].Category, "Word") {
// 			count++
// 			if count < 100 {
// 				riddles[i].Category = "General"
// 			} else {
// 				riddles[i].Category = "Logic"
// 			}
// 		}
// 		if count < 50 {
// 			riddles[i].Difficulty = "Hard"
// 			riddles[i].Points = 15
// 		}
// 		if count > 70 && count < 120 {
// 			riddles[i].Difficulty = "Easy"
// 		}
// 		if count > 121 {
// 			riddles[i].Difficulty = "Medium"
// 		}

// 		if riddles[i].Difficulty == "Easy" {
// 			riddles[i].Points = 5
// 		}
// 		if riddles[i].Difficulty == "Medium" {
// 			riddles[i].Points = 10
// 		}
// 		if riddles[i].Difficulty == "Hard" {
// 			riddles[i].Points = 15
// 		}

// 		err := db.Create(&riddles[i]).Error
// 		if err != nil {
// 			// Log the error and stop processing further
// 			fmt.Printf("Failed to create riddle at index %d: %v\n", i, err)
// 			break
// 		}
// 		time.Sleep(time.Second * 3)
// 	}

// }
