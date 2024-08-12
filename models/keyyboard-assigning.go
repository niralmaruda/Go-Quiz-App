package models

import (
	"strconv"

	"github.com/eiannone/keyboard"
)

type HandleKey struct {
	strategy map[keyboard.Key]Executable
}

// Initialize the strategy map
func NewHandleKey() *HandleKey {
	return &HandleKey{
		strategy: map[keyboard.Key]Executable{
			keyboard.KeyEnter:      KeyEnter{},
			keyboard.KeyArrowUp:    KeyArrowUp{},
			keyboard.KeyArrowDown:  KeyArrowDown{},
			keyboard.KeyArrowLeft:  KeyArrowLeft{},
			keyboard.KeyArrowRight: KeyArrowRight{},
		},
	}
}

// Execute the strategy based on the key pressed
func (h *HandleKey) Execute(key keyboard.Key, selectedOption *int) string {
	if executable, exists := h.strategy[key]; exists {
		return executable.Execute(selectedOption)
	}
	return strconv.Itoa(*selectedOption) // Default return value if key is not mapped
}
