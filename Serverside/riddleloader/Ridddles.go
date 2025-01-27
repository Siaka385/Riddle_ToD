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

var FilteredRiddles []Riddle
var RiddleChoicesMap map[int][]Choice = make(map[int][]Choice)
var RiddleHintsMap map[int][]Hint = make(map[int][]Hint)

// LoadRiddlesFromDB retrieves riddles based on the selected category, generates choices and hints, and writes the riddles to a JSON file.
func LoadRiddlesFromDB(w http.ResponseWriter, r *http.Request, category string, db *gorm.DB) {
	// Retrieve riddles from the database
	err := db.Model(&database.Riddle{}).
		Select("ID, Question, Answer, Explanation, Category, Difficulty, Points").
		Where("category = ?", category).
		Find(&FilteredRiddles).Error

	if err != nil {
		fmt.Println("Error retrieving riddles:", err)
		os.Exit(1)
	}
	// Write the filtered riddles to a JSON file
	WriteRiddlesToFile(w, r)

	// Generate choices and hints for the riddles
	GenerateRiddleChoices(db)
	//fmt.Println(RiddleChoicesMap)
	GenerateRiddleHints(db)
	//fmt.Println(RiddleHintsMap)
}

// WriteRiddlesToFile writes the filtered riddles to a JSON file.
func WriteRiddlesToFile(w http.ResponseWriter, r *http.Request) {
	if len(FilteredRiddles) == 0 {
		fmt.Println("No riddles to save to file.")
		os.Exit(1)
	}

	file, err := os.Create("FilteredRiddles.json")
	if err != nil {
		fmt.Println("Error creating the JSON file.")
		os.Exit(1)
	}
	defer file.Close()

	// Serialize the riddles to JSON format
	jsonData, err := json.MarshalIndent(FilteredRiddles, "", "  ")
	if err != nil {
		fmt.Println("Error converting riddles to JSON.")
		os.Exit(1)
	}

	// Write the JSON data to the file
	err = os.WriteFile(file.Name(), jsonData, 0o644)
	if err != nil {
		fmt.Println("Error writing to JSON file.")
		os.Exit(1)
	}
}

// GenerateRiddleChoices populates the RiddleChoicesMap with choices for each riddle.
func GenerateRiddleChoices(db *gorm.DB) {
	for i := 0; i < len(FilteredRiddles); i++ {
		var choices []Choice // Reset the slice for each riddle

		// Retrieve choices for the current riddle
		err := db.Model(&database.Choice{}).
			Select("ID, Text").
			Where("riddle_id = ?", FilteredRiddles[i].ID).
			Find(&choices).Error

		if err != nil {
			fmt.Println("Error retrieving choices for riddle ID", FilteredRiddles[i].ID, ":", err)
			continue
		}

		// Store the choices in the map
		RiddleChoicesMap[FilteredRiddles[i].ID] = choices
	}
}

// GenerateRiddleHints populates the RiddleHintsMap with hints for each riddle.
func GenerateRiddleHints(db *gorm.DB) {
	for i := 0; i < len(FilteredRiddles); i++ {
		var hints []Hint // Reset the slice for each riddle

		// Retrieve hints for the current riddle
		err := db.Model(&database.Hint{}).
			Select("ID, Text").
			Where("riddle_id = ?", FilteredRiddles[i].ID).
			Find(&hints).Error

		if err != nil {
			fmt.Println("Error retrieving hints for riddle ID", FilteredRiddles[i].ID, ":", err)
			continue
		}

		// Store the hints in the map
		RiddleHintsMap[FilteredRiddles[i].ID] = hints
	}
}
