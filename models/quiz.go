package models

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/eiannone/keyboard"
)

type Quiz struct {
	question        string
	options         []Option
	isMultiSelected bool
}

type Option struct {
	option  string
	correct bool
}

var quiz []Quiz

func Init() []Quiz {
	quiz = []Quiz{{
		question: "How are you",
		options: []Option{
			{option: "Good", correct: true},
			{option: "Bad", correct: false},
			{option: "Okay", correct: false},
			{option: "Excellent", correct: false},
		},
		isMultiSelected: false,
	},
		{
			question: "What is the closest planet to Earth?",
			options: []Option{
				{option: "Venus", correct: false},
				{option: "Mars", correct: true},
				{option: "Jupiter", correct: false},
				{option: "Saturn", correct: false},
			},
			isMultiSelected: false,
		},
	}

	return quiz
}

func (q *Quiz) PrintQuestion() {
	fmt.Println(q.question)
}

func (q *Quiz) PrintOptions(selectedOption *int) {
	for i, option := range q.options {
		if selectedOption != nil && *selectedOption == i+1 {
			fmt.Printf("> %d. %s\n", i+1, option.option)
		} else {
			fmt.Printf("%d. %s\n", i+1, option.option)
		}
	}
}

func (q *Quiz) CheckAnswer(input string) {
	correct := false
	if !q.isMultiSelected {
		for i, option := range q.options {
			if input == fmt.Sprint(i+1) && option.correct {
				correct = true
				break
			}
		}
	} else {
		// Additional logic for checkbox type questions
		fmt.Print("No implementation for checkbox yet.")
	}

	if correct {
		fmt.Println("Correct!")
	} else {
		fmt.Println("incorrect!")
	}
}

func (q *Quiz) ReadInput() string {
	handleKey := NewHandleKey()
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	selectedOption := 1

	for {
		_, key, _ := keyboard.GetKey()
		q.ClearTerminal()
		q.PrintQuestion()
		q.PrintOptions(&selectedOption)
		input := handleKey.Execute(key, &selectedOption)
		fmt.Println("Enter your answer: " + input)
		if key == keyboard.KeyEnter {
			q.CheckAnswer(input)
			return input
		}
	}
}

func (q *Quiz) ClearTerminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
