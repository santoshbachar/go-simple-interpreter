package lexer

import (
	"strconv"
)

func IsDigit(char string) bool {
	if _, err := strconv.Atoi(char); err == nil {
		return true
	}
	return false
}

func IsCharADigit(char byte) bool {
	digits := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for _, val := range digits {
		if char == val {
			return true
		}
	}
	return false
}
