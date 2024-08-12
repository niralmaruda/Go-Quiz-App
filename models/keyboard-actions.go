package models

import (
	"strconv"
)

type Executable interface {
	Execute(selectedOption *int) string
}

type KeyEnter struct{}
type KeyArrowUp struct{}
type KeyArrowDown struct{}
type KeyArrowLeft struct{}
type KeyArrowRight struct{}

func (k KeyEnter) Execute(selectedOption *int) string {
	return strconv.Itoa(*selectedOption)
}

func (k KeyArrowUp) Execute(selectedOption *int) string {
	if *selectedOption > 1 {
		*selectedOption -= 1
	}
	return strconv.Itoa(*selectedOption)
}

func (k KeyArrowDown) Execute(selectedOption *int) string {
	if *selectedOption < 4 {
		*selectedOption += 1
	}
	return strconv.Itoa(*selectedOption)
}

func (k KeyArrowLeft) Execute(selectedOption *int) string {
	*selectedOption = 0
	return strconv.Itoa(*selectedOption)
}

func (k KeyArrowRight) Execute(selectedOption *int) string {
	return strconv.Itoa(*selectedOption)
}
