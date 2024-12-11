package Serverside

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Riddle struct {
	Question   string   `json:"question"`   // The riddle question
	Category   string   `json:"category"`   // The category of the riddle (e.g., General, Logic, Mathematics, Word Riddle)
	Choices    []string `json:"choices"`    // Multiple choice options for the riddle
	Difficulty string   `json:"difficulty"` // The difficulty level of the riddle (easy, medium, hard)
	Level      int      `json:"level"`      // The level of the riddle
}

var SelectedRiddles []Riddle

// LoadRiddle loads riddles from a JSON file, filters them based on selected categories, and writes filtered riddles to a new file.
func LoadRiddle(w http.ResponseWriter, r *http.Request, selectedCategories RiddleCategoteries) {
	// Open the riddle JSON file
	file, err := os.Open("riddle.json")
	if err != nil {
		fmt.Println("Error opening the file")
		os.Exit(1)
	}
	defer file.Close()

	SelectedRiddles = []Riddle{}

	jsonDecoder := json.NewDecoder(file)

	// Read the opening JSON bracket
	jsonDecoder.Token()

	var currentRiddle Riddle

	// Process each JSON object in the array
	for jsonDecoder.More() {
		err := jsonDecoder.Decode(&currentRiddle)
		if err != nil {
			fmt.Println("Error decoding JSON")
			os.Exit(1)
		}

		FilterByCategory(selectedCategories, currentRiddle)
	}

	jsonDecoder.Token()

	// Write filtered riddles to a new file
	WriteFilteredRiddles(w, r)
}

// FilterByCategory determines the category selected and filters the riddles accordingly.
func FilterByCategory(category RiddleCategoteries, riddle Riddle) {
	// Check the selected category and pass the riddle to the filtering function
	if category.General {
		FilterCategory(riddle, "General") // Filter riddles in the "General" category
	} else if category.Logic {
		FilterCategory(riddle, "Logic") // Filter riddles in the "Logic" category
	} else if category.Mathematics {
		FilterCategory(riddle, "Mathematics") // Filter riddles in the "Mathematics" category
	} else if category.Wordriddle {
		FilterCategory(riddle, "Word Riddle") // Filter riddles in the "Word Riddle" category
	} else {
		// Handle unknown categories by exiting with an error message
		fmt.Println("Unknown category")
		os.Exit(1)
	}
}

// FilterCategory appends a riddle to the selected riddles list if its category matches the specified one.
func FilterCategory(riddle Riddle, category string) {
	if riddle.Category == category {
		SelectedRiddles = append(SelectedRiddles, riddle)
	}
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
