package models

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Quiz struct {
	question string
	options  []Option
	quesType int
}

type Option struct {
	option  string
	correct bool
}

const (
	radio = iota
	checkbox
)

const (
	w = iota
	s
	a
	d
)

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
		quesType: radio,
	}}

	return quiz
}

func (q Quiz) PrintQuestion() {
	fmt.Println(q.question)
}

func (q Quiz) PrintOptions() {
	for i, option := range q.options {
		fmt.Printf("%d. %s\n", i+1, option.option)
	}
}

func (q Quiz) CheckAnswer(input string) {
	correct := false
	if q.quesType == radio {
		for i, option := range q.options {
			if input == fmt.Sprint(i+1) && option.correct {
				correct = true
				break
			}
		}
	} else if q.quesType == checkbox {
		// Additional logic for checkbox type questions
		fmt.Print("No implementation for checkbox yet.")
	}

	if correct {
		fmt.Println("Correct!")
	} else {
		fmt.Println("incorrect!")
	}
}

func (q Quiz) ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your choice: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}
