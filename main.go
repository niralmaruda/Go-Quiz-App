package main

import (
	"fmt"

	"main.go/models"
)

func main() {
	fmt.Println("Welcome to GO!")
	quiz := models.Init()

	for _, q := range quiz {
		printQuestions(q)
	}
}

func printQuestions(quiz models.Quiz) {

	quiz.ReadInput()

}
