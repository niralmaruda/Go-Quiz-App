package models

import (
	"fmt"
	"log"

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
	}}

	return quiz
}

func (q *Quiz) PrintQuestion() {
	fmt.Println(q.question)
}

func (q *Quiz) PrintOptions() {
	for i, option := range q.options {
		fmt.Printf("%d. %s\n", i+1, option.option)
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
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	selectedOption := 0
	var userInput string

	for {
		ch, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch key {
		case keyboard.KeyArrowUp:
			if selectedOption > 0 {
				selectedOption--
			}
		case keyboard.KeyArrowDown:
			if selectedOption < len(q.options)-1 {
				selectedOption++
			}
		case keyboard.KeyArrowRight:
			userInput = fmt.Sprint(selectedOption + 1)
		case keyboard.KeyArrowLeft:
			userInput = "0"
		case keyboard.KeyEnter:
			return userInput
		}

		if ch == 'q' || ch == 'Q' {
			return "0"
		}
	}
}

func Exec() {
	var keys = map[keyboard.Key]Executable{
		keyboard.KeyEnter:      KeyEnter{},
		keyboard.KeyArrowUp:    KeyArrowUp{},
		keyboard.KeyArrowDown:  KeyArrowDown{},
		keyboard.KeyArrowLeft:  KeyArrowLeft{},
		keyboard.KeyArrowRight: KeyArrowRight{},
	}
}

type Executable interface {
	Execute()
}

type KeyEnter struct{}
type KeyArrowUp struct{}
type KeyArrowDown struct{}
type KeyArrowLeft struct{}
type KeyArrowRight struct{}

func (k KeyEnter) Execute() {
	fmt.Println("KeyEnter")
}

func (k KeyArrowUp) Execute() {
	fmt.Println("KeyArrowUp")
}

func (k KeyArrowDown) Execute() {
	fmt.Println("KeyArrowDown")
}

func (k KeyArrowLeft) Execute() {
	fmt.Println("KeyArrowLeft")
}

func (k KeyArrowRight) Execute() {
	fmt.Println("KeyArrowRight")
}
