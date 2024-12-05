package Serverside

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Riddle struct {
	Question  string   `json:"qustion"`    // riddle question
	Category  string   `json:"category"`   // riddle category (logic,word riddle,mathematics,logic)
	Choices   []string `json:"choices"`    // multiple choices for the riddle
	Difficult string   `json:"difficulty"` // difficulty of the question(easy,medium,hard)
	Level     int      `json:"level"`      // level of   the qurstion
}

func LoadRiddle(w http.ResponseWriter, r *http.Request) {
	riddles := []Riddle{
		{
			Question:  "I speak without a mouth and hear without ears. I have no body, but I come alive with the wind. What am I?",
			Category:  "Logic",
			Choices:   []string{"Echo", "Shadow", "Wind", "Sound"},
			Difficult: "Medium",
			Level:     1,
		},
		{
			Question:  "The more of me you take, the more you leave behind. What am I?",
			Category:  "Logic",
			Choices:   []string{"Footsteps", "Memories", "Time", "Shadow"},
			Difficult: "Medium",
			Level:     1,
		},
		{
			Question:  "What has keys but can't open locks?",
			Category:  "Word Riddle",
			Choices:   []string{"Piano", "Keyboard", "Map", "Book"},
			Difficult: "Easy",
			Level:     1,
		},
		{
			Question:  "What comes once in a minute, twice in a moment, but never in a thousand years?",
			Category:  "Word Riddle",
			Choices:   []string{"The letter 'M'", "Time", "Infinity", "Silence"},
			Difficult: "Medium",
			Level:     2,
		},
		{
			Question:  "What number comes next in the sequence: 1, 11, 21, 1211, 111221?",
			Category:  "Mathematics",
			Choices:   []string{"312211", "111211", "131122", "211221"},
			Difficult: "Hard",
			Level:     3,
		},
		{
			Question:  "I’m tall when I’m young, and I’m short when I’m old. What am I?",
			Category:  "Logic",
			Choices:   []string{"Candle", "Tree", "Shadow", "Mountain"},
			Difficult: "Easy",
			Level:     1,
		},
		{
			Question:  "What is always in front of you but can’t be seen?",
			Category:  "Logic",
			Choices:   []string{"Future", "Air", "Dream", "Reflection"},
			Difficult: "Medium",
			Level:     1,
		},
		{
			Question:  "What has one eye but can’t see?",
			Category:  "Word Riddle",
			Choices:   []string{"Needle", "Cyclops", "Storm", "Clock"},
			Difficult: "Easy",
			Level:     1,
		},
		{
			Question:  "If two’s company and three’s a crowd, what are four and five?",
			Category:  "Word Riddle",
			Choices:   []string{"Nine", "Party", "Friends", "Math Problem"},
			Difficult: "Medium",
			Level:     2,
		},
		{
			Question:  "If you have me, you want to share me. If you share me, you don’t have me. What am I?",
			Category:  "Logic",
			Choices:   []string{"Secret", "Happiness", "Silence", "Gift"},
			Difficult: "Hard",
			Level:     3,
		},
		{
			Question:  "I’m found in the sky, but you can also hold me in your hands. What am I?",
			Category:  "Logic",
			Choices:   []string{"Cloud", "Star", "Water", "Dream"},
			Difficult: "Medium",
			Level:     2,
		},
		{
			Question:  "What has four wheels and flies?",
			Category:  "Word Riddle",
			Choices:   []string{"Garbage Truck", "Airplane", "Helicopter", "Race Car"},
			Difficult: "Easy",
			Level:     1,
		},
		{
			Question:  "The more you take away, the bigger I get. What am I?",
			Category:  "Logic",
			Choices:   []string{"Hole", "Shadow", "Silence", "Cloud"},
			Difficult: "Medium",
			Level:     2,
		},
		{
			Question:  "I am always hungry, I must always be fed. The finger I touch, will soon turn red. What am I?",
			Category:  "Logic",
			Choices:   []string{"Fire", "Lava", "Flame", "Sun"},
			Difficult: "Medium",
			Level:     2,
		},
		{
			Question:  "What is so fragile that saying its name breaks it?",
			Category:  "Logic",
			Choices:   []string{"Silence", "Glass", "Promise", "Shadow"},
			Difficult: "Medium",
			Level:     2,
		},
		{
			Question:  "What has roots as nobody sees, is taller than trees, up, up it goes, and yet it never grows?",
			Category:  "Logic",
			Choices:   []string{"Mountain", "Tree", "Cloud", "Building"},
			Difficult: "Hard",
			Level:     3,
		},
		{
			Question:  "I can be cracked, made, told, and played. What am I?",
			Category:  "Word Riddle",
			Choices:   []string{"Joke", "Puzzle", "Song", "Game"},
			Difficult: "Easy",
			Level:     1,
		},
		{
			Question:  "The person who makes it sells it. The person who buys it never uses it. The person who uses it never knows they’re using it. What is it?",
			Category:  "Logic",
			Choices:   []string{"Coffin", "Clock", "Medicine", "Gift"},
			Difficult: "Hard",
			Level:     3,
		},
		{
			Question:  "What has a head, a tail, is brown, and has no legs?",
			Category:  "Logic",
			Choices:   []string{"Coin", "Worm", "Snail", "Stick"},
			Difficult: "Easy",
			Level:     1,
		},
		{
			Question:  "What begins with an ‘E’, but only contains one letter?",
			Category:  "Word Riddle",
			Choices:   []string{"Envelope", "Eagle", "Edge", "Eel"},
			Difficult: "Easy",
			Level:     1,
		},
	}

	json, err := json.MarshalIndent(riddles, "", " ")
	if err != nil {
		fmt.Println("ERROR ON STRUCT TO JSON")
		os.Exit(1)
	}

	file, err := os.Create("riddle.json")
	if err != nil {
		fmt.Println("Error Creating file")
		os.Exit(1)
	}
	defer file.Close()
	if err := os.WriteFile("riddle.json", json, 0o064); err != nil {
		fmt.Println("Error Creating file")
		os.Exit(1)
	}
}
