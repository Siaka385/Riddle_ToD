package addriddles

import (
	"fmt"
	"os"

	"gorm.io/gorm"

	"Riddle_ToD/Serverside/database"
)


func Addriddle(db *gorm.DB) {

	riddles := []database.Riddle{
		{
			ID:         1,
			Question:   "What has to be broken before you can use it?",
			Answer:     "An egg",
			Category:   "Logic",
			Difficulty: "Easy",
			Points:     5,
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
		{
			ID:         2,
			Question:   "I’m tall when I’m young, and I’m short when I’m old. What am I?",
			Answer:     "A candle",
			Category:   "Logic",
			Difficulty: "Easy",
			Points:     5,
			Choices: []database.Choice{
				{ID: 5, RiddleID: 2, Text: "A candle"},
				{ID: 6, RiddleID: 2, Text: "A pencil"},
				{ID: 7, RiddleID: 2, Text: "A tree"},
				{ID: 8, RiddleID: 2, Text: "A shadow"},
			},
			Hints: []database.Hint{
				{ID: 3, RiddleID: 2, Text: "It gives light."},
				{ID: 4, RiddleID: 2, Text: "It melts as it’s used."},
			},
		},
		{
			ID:         3,
			Question:   "What has a heart that doesn’t beat?",
			Answer:     "An artichoke",
			Category:   "General",
			Difficulty: "Medium",
			Points:     10,
			Choices: []database.Choice{
				{ID: 9, RiddleID: 3, Text: "An artichoke"},
				{ID: 10, RiddleID: 3, Text: "A clock"},
				{ID: 11, RiddleID: 3, Text: "A stone"},
				{ID: 12, RiddleID: 3, Text: "A statue"},
			},
			Hints: []database.Hint{
				{ID: 5, RiddleID: 3, Text: "It’s a type of vegetable."},
				{ID: 6, RiddleID: 3, Text: "It’s often eaten in salads or dips."},
			},
		},
		{
			ID:         4,
			Question:   "I’m tall when I’m young, and I’m short when I’m old. What am I?",
			Answer:     "A candle",
			Category:   "Logic",
			Difficulty: "Easy",
			Points:     5,
			Choices: []database.Choice{
				{ID: 13, RiddleID: 4, Text: "A tree"},
				{ID: 14, RiddleID: 4, Text: "A candle"},
				{ID: 15, RiddleID: 4, Text: "A pencil"},
				{ID: 16, RiddleID: 4, Text: "A shadow"},
			},
			Hints: []database.Hint{
				{ID: 7, RiddleID: 4, Text: "It is made of wax."},
				{ID: 8, RiddleID: 4, Text: "It gives light."},
			},
		},
		{
			ID:         5,
			Question:   "What has a heart that doesn’t beat?",
			Answer:     "An artichoke",
			Category:   "General",
			Difficulty: "Medium",
			Points:     10,
			Choices: []database.Choice{
				{ID: 17, RiddleID: 5, Text: "A robot"},
				{ID: 18, RiddleID: 5, Text: "An artichoke"},
				{ID: 19, RiddleID: 5, Text: "A clock"},
				{ID: 20, RiddleID: 5, Text: "A stone"},
			},
			Hints: []database.Hint{
				{ID: 9, RiddleID: 5, Text: "It is a vegetable."},
				{ID: 10, RiddleID: 5, Text: "You eat its leaves."},
			},
		},
		{
			ID:         6,
			Question:   "What has keys but can’t open locks?",
			Answer:     "A piano",
			Category:   "General",
			Difficulty: "Easy",
			Points:     5,
			Choices: []database.Choice{
				{ID: 21, RiddleID: 6, Text: "A keyboard"},
				{ID: 22, RiddleID: 6, Text: "A piano"},
				{ID: 23, RiddleID: 6, Text: "A lockbox"},
				{ID: 24, RiddleID: 6, Text: "A treasure chest"},
			},
			Hints: []database.Hint{
				{ID: 11, RiddleID: 6, Text: "It makes music."},
				{ID: 12, RiddleID: 6, Text: "It has black and white keys."},
			},
		},
		{
			ID:         7,
			Question:   "The more you take, the more you leave behind. What am I?",
			Answer:     "Footsteps",
			Category:   "Logic",
			Difficulty: "Medium",
			Points:     10,
			Choices: []database.Choice{
				{ID: 25, RiddleID: 7, Text: "Sand"},
				{ID: 26, RiddleID: 7, Text: "Footsteps"},
				{ID: 27, RiddleID: 7, Text: "Time"},
				{ID: 28, RiddleID: 7, Text: "Shadows"},
			},
			Hints: []database.Hint{
				{ID: 13, RiddleID: 7, Text: "It’s related to walking."},
				{ID: 14, RiddleID: 7, Text: "It is left on the ground."},
			},
		},
		{
			ID:         8,
			Question:   "What has a face and two hands but no arms or legs?",
			Answer:     "A clock",
			Category:   "General",
			Difficulty: "Easy",
			Points:     5,
			Choices: []database.Choice{
				{ID: 29, RiddleID: 8, Text: "A clock"},
				{ID: 30, RiddleID: 8, Text: "A robot"},
				{ID: 31, RiddleID: 8, Text: "A compass"},
				{ID: 32, RiddleID: 8, Text: "A doll"},
			},
			Hints: []database.Hint{
				{ID: 15, RiddleID: 8, Text: "It tells time."},
				{ID: 16, RiddleID: 8, Text: "It’s found on walls or wrists."},
			},
		},
		{
			ID:         9,
			Question:   "I’m light as a feather, yet the strongest person can’t hold me for more than a minute. What am I?",
			Answer:     "Your breath",
			Category:   "Logic",
			Difficulty: "Medium",
			Points:     10,
			Choices: []database.Choice{
				{ID: 33, RiddleID: 9, Text: "A balloon"},
				{ID: 34, RiddleID: 9, Text: "Your breath"},
				{ID: 35, RiddleID: 9, Text: "A thought"},
				{ID: 36, RiddleID: 9, Text: "A cloud"},
			},
			Hints: []database.Hint{
				{ID: 17, RiddleID: 9, Text: "It’s part of your body."},
				{ID: 18, RiddleID: 9, Text: "You need it to survive."},
			},
		},
		{
			ID:         10,
			Question:   "I’m always on my way but never arrive. What am I?",
			Answer:     "Tomorrow",
			Category:   "Logic",
			Difficulty: "Hard",
			Points:     15,
			Choices: []database.Choice{
				{ID: 37, RiddleID: 10, Text: "A train"},
				{ID: 38, RiddleID: 10, Text: "Tomorrow"},
				{ID: 39, RiddleID: 10, Text: "The future"},
				{ID: 40, RiddleID: 10, Text: "A comet"},
			},
			Hints: []database.Hint{
				{ID: 19, RiddleID: 10, Text: "It’s part of time."},
				{ID: 20, RiddleID: 10, Text: "It never arrives today."},
			},
		},
		{
			ID:         11,
			Question:   "What gets wetter the more it dries?",
			Answer:     "A towel",
			Category:   "Logic",
			Difficulty: "Easy",
			Points:     5,
			Choices: []database.Choice{
				{ID: 41, RiddleID: 11, Text: "A sponge"},
				{ID: 42, RiddleID: 11, Text: "A towel"},
				{ID: 43, RiddleID: 11, Text: "Rain"},
				{ID: 44, RiddleID: 11, Text: "A cloud"},
			},
			Hints: []database.Hint{
				{ID: 21, RiddleID: 11, Text: "You use it after a shower."},
				{ID: 22, RiddleID: 11, Text: "It absorbs water."},
			},
		},
		{
			ID:         12,
			Question:   "What has a head, a tail, is brown, and has no legs?",
			Answer:     "A penny",
			Category:   "General",
			Difficulty: "Easy",
			Points:     5,
			Choices: []database.Choice{
				{ID: 45, RiddleID: 12, Text: "A coin"},
				{ID: 46, RiddleID: 12, Text: "A penny"},
				{ID: 47, RiddleID: 12, Text: "A snake"},
				{ID: 48, RiddleID: 12, Text: "A worm"},
			},
			Hints: []database.Hint{
				{ID: 23, RiddleID: 12, Text: "It is money."},
				{ID: 24, RiddleID: 12, Text: "It’s often copper-colored."},
			},
		},
		{
			ID:         13,
			Question:   "What comes once in a minute, twice in a moment, but never in a thousand years?",
			Answer:     "The letter M",
			Category:   "Word",
			Difficulty: "Medium",
			Points:     10,
			Choices: []database.Choice{
				{ID: 49, RiddleID: 13, Text: "The letter M"},
				{ID: 50, RiddleID: 13, Text: "Time"},
				{ID: 51, RiddleID: 13, Text: "A second"},
				{ID: 52, RiddleID: 13, Text: "Infinity"},
			},
			Hints: []database.Hint{
				{ID: 25, RiddleID: 13, Text: "It’s part of a word."},
				{ID: 26, RiddleID: 13, Text: "It’s a letter in the alphabet."},
			},
		},
		{
			ID:         14,
			Question:   "I have branches, but no fruit, trunk, or leaves. What am I?",
			Answer:     "A bank",
			Category:   "Logic",
			Difficulty: "Medium",
			Points:     10,
			Choices: []database.Choice{
				{ID: 53, RiddleID: 14, Text: "A river"},
				{ID: 54, RiddleID: 14, Text: "A bank"},
				{ID: 55, RiddleID: 14, Text: "A tree"},
				{ID: 56, RiddleID: 14, Text: "A map"},
			},
			Hints: []database.Hint{
				{ID: 27, RiddleID: 14, Text: "It involves money."},
				{ID: 28, RiddleID: 14, Text: "It has physical locations."},
			},
		},
		{
			ID:         15,
			Question:   "The more you take away from me, the bigger I get. What am I?",
			Answer:     "A hole",
			Category:   "Logic",
			Difficulty: "Medium",
			Points:     10,
			Choices: []database.Choice{
				{ID: 57, RiddleID: 15, Text: "A hole"},
				{ID: 58, RiddleID: 15, Text: "A crater"},
				{ID: 59, RiddleID: 15, Text: "A shadow"},
				{ID: 60, RiddleID: 15, Text: "A vacuum"},
			},
			Hints: []database.Hint{
				{ID: 29, RiddleID: 15, Text: "It’s an empty space."},
				{ID: 30, RiddleID: 15, Text: "You dig it to make it bigger."},
			},
		},
		{
			ID:         16,
			Question:   "What has an eye but cannot see?",
			Answer:     "A needle",
			Category:   "General",
			Difficulty: "Easy",
			Points:     5,
			Choices: []database.Choice{
				{ID: 61, RiddleID: 16, Text: "A needle"},
				{ID: 62, RiddleID: 16, Text: "A storm"},
				{ID: 63, RiddleID: 16, Text: "A potato"},
				{ID: 64, RiddleID: 16, Text: "A blindfold"},
			},
			Hints: []database.Hint{
				{ID: 31, RiddleID: 16, Text: "It’s used in sewing."},
				{ID: 32, RiddleID: 16, Text: "Thread passes through it."},
			},
		},
		{
			ID:         17,
			Question:   "I’m not alive, but I can grow. I don’t have lungs, but I need air. What am I?",
			Answer:     "Fire",
			Category:   "Logic",
			Difficulty: "Medium",
			Points:     10,
			Choices: []database.Choice{
				{ID: 65, RiddleID: 17, Text: "Fire"},
				{ID: 66, RiddleID: 17, Text: "A plant"},
				{ID: 67, RiddleID: 17, Text: "A balloon"},
				{ID: 68, RiddleID: 17, Text: "A shadow"},
			},
			Hints: []database.Hint{
				{ID: 33, RiddleID: 17, Text: "It’s hot and bright."},
				{ID: 34, RiddleID: 17, Text: "It consumes fuel to exist."},
			},
		},
		{
			ID:         18,
			Question:   "The more you have of me, the less you see. What am I?",
			Answer:     "Darkness",
			Category:   "Logic",
			Difficulty: "Medium",
			Points:     10,
			Choices: []database.Choice{
				{ID: 69, RiddleID: 18, Text: "Fog"},
				{ID: 70, RiddleID: 18, Text: "Darkness"},
				{ID: 71, RiddleID: 18, Text: "A blindfold"},
				{ID: 72, RiddleID: 18, Text: "A shadow"},
			},
			Hints: []database.Hint{
				{ID: 35, RiddleID: 18, Text: "It happens at night."},
				{ID: 36, RiddleID: 18, Text: "It’s the absence of light."},
			},
		},
		{
			ID:         19,
			Question:   "What has a neck but no head?",
			Answer:     "A bottle",
			Category:   "General",
			Difficulty: "Easy",
			Points:     5,
			Choices: []database.Choice{
				{ID: 73, RiddleID: 19, Text: "A bottle"},
				{ID: 74, RiddleID: 19, Text: "A shirt"},
				{ID: 75, RiddleID: 19, Text: "A vase"},
				{ID: 76, RiddleID: 19, Text: "A guitar"},
			},
			Hints: []database.Hint{
				{ID: 37, RiddleID: 19, Text: "It holds liquids."},
				{ID: 38, RiddleID: 19, Text: "It’s made of glass or plastic."},
			},
		},
		{
			ID:         20,
			Question:   "I speak without a mouth and hear without ears. I have no body, but I come alive with the wind. What am I?",
			Answer:     "An echo",
			Category:   "Logic",
			Difficulty: "Hard",
			Points:     15,
			Choices: []database.Choice{
				{ID: 77, RiddleID: 20, Text: "An echo"},
				{ID: 78, RiddleID: 20, Text: "A ghost"},
				{ID: 79, RiddleID: 20, Text: "The wind"},
				{ID: 80, RiddleID: 20, Text: "A shadow"},
			},
			Hints: []database.Hint{
				{ID: 39, RiddleID: 20, Text: "It’s a sound."},
				{ID: 40, RiddleID: 20, Text: "It repeats what you say."},
			},
		},
	}

}
