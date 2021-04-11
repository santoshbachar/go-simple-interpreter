package lexer

import "testing"

func TestisCharADigit(t *testing.T) {
	isDigit := isCharADigit('5')

	if isDigit == false {
		t.Errorf("bool is incorrent")
	}
}
