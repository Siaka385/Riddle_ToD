package riddleloader

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"gorm.io/gorm"

	"Riddle_ToD/Serverside/database"
)

type Riddle struct {
	ID          int    `json:"id"`
	Question    string `json:"question"`
	Answer      string `json:"answer"`
	Explanation string `json:"explanation"`
	Category    string `json:"category"`
	Difficulty  string `json:"difficulty"`
	Points      int    `json:"points"`
}
type Choice struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type Hint struct {
	ID       int    `json:"id"`
	RiddleID int    `json:"riddle_id"`
	Text     string `json:"text"`
}

var SelectedRiddles []Riddle
var RiddleChoices map[int][]Choice = make(map[int][]Choice)
var RiddleHint map[int][]Hint = make(map[int][]Hint)

// LoadRiddle loads riddles from a JSON file, filters them based on selected categories, and writes filtered riddles to a new file.
func LoadRiddle(w http.ResponseWriter, r *http.Request, selectedCategories string, db *gorm.DB) {

	err := db.Model(&database.Riddle{}).
		Select("ID,Question,Answer,Explanation,Category,Difficulty,Points").
		Where("category = ?", selectedCategories).
		Find(&SelectedRiddles).Error

	if err != nil {
		fmt.Println("Error retrieving riddles:", err)
	} else {
		fmt.Println("Retrieved riddles:", SelectedRiddles[0])
	}

	// Write filtered riddles to a new file
	WriteFilteredRiddles(w, r)
	GenerateChoices(db)
	fmt.Println(RiddleChoices)
	GenerateHints(db)
	fmt.Println(RiddleHint)
}

// WriteFilteredRiddles writes the filtered riddles to a new JSON file.
func WriteFilteredRiddles(w http.ResponseWriter, r *http.Request) {
	if len(SelectedRiddles) == 0 {
		fmt.Println("No riddles to select")
		os.Exit(1)
	}

	file, err := os.Create("FilteredRiddles.json")
	if err != nil {
		fmt.Println("Error creating the JSON file")
		os.Exit(1)
	}
	defer file.Close()

	// Serialize the riddles to JSON format
	jsonData, err := json.MarshalIndent(SelectedRiddles, "", "  ")
	if err != nil {
		fmt.Println("Error converting riddles to JSON")
		os.Exit(1)
	}

	// Write the JSON data to the file
	err = os.WriteFile(file.Name(), jsonData, 0o644)
	if err != nil {
		fmt.Println("Error writing to JSON file")
		os.Exit(1)
	}
}

func GenerateChoices(db *gorm.DB) {
	// Loop through the selected riddles
	for i := 0; i < len(SelectedRiddles); i++ {
		var riddleChoices []Choice // Reset the slice to avoid appending from previous iterations

		// Retrieve the choices for the current riddle
		err := db.Model(&database.Choice{}).
			Select("ID, Text").
			Where("riddle_id = ?", SelectedRiddles[i].ID). // Note that riddle_id is the foreign key in the database
			Find(&riddleChoices).Error
		//	time.Sleep(3 * time.Second)
		if err != nil {
			fmt.Println("ERROR RETRIEVING THE CHOICES:", err)
			continue // Skip to the next riddle if an error occurs
		}

		// Store the retrieved choices in the map
		RiddleChoices[SelectedRiddles[i].ID] = riddleChoices
	}
}

func GenerateHints(db *gorm.DB) {
	// Loop through the selected riddles
	for i := 0; i < len(SelectedRiddles); i++ {
		var riddleHints []Hint // Reset the slice to avoid appending from previous iterations

		// Retrieve the hints for the current riddle
		err := db.Model(&Hint{}).
			Select("ID, Text").
			Where("riddle_id = ?", SelectedRiddles[i].ID). // Note that riddle_id is the foreign key in the database
			Find(&riddleHints).Error

		if err != nil {
			fmt.Println("ERROR RETRIEVING THE HINTS:", err)
			continue // Skip to the next riddle if an error occurs
		}

		// Store the retrieved hints in the map
		RiddleHint[SelectedRiddles[i].ID] = riddleHints
	}
}
