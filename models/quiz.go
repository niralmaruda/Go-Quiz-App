package models

type Quiz struct {
	Question string
	Options  []Option
}

type Option struct {
	Option  string
	Correct bool
}

type QuesType struct {
	Radio    int
	Checkbox int
}

const (
	Radio = iota
	Checkbox
)
