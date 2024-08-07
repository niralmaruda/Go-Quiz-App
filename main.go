package main

import (
	"fmt"

	"main.go/models"
)

func main() {
	fmt.Println("Welcome to GO!")
	ques := models.Quiz{
		Question: "How are you",
		Options: []models.Option{
			{Option: "Good", Correct: true},
			{Option: "Bad", Correct: false},
			{Option: "Okay", Correct: false},
			{Option: "Excellent", Correct: false},
		},
	}
	fmt.Print(ques)
}
