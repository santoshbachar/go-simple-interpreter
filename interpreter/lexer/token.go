package lexer

import "fmt"

// TokenType for our interpreter
type TokenType string

const (
	// INTEGER value
	INTEGER TokenType = "INTEGER"
	// SUM upon calculation
	SUM = "SUM"
	// EOF just return now
	EOF = "EOF"
)

// TOKEN to return after parse by Interpreter to eat
type TOKEN struct {
	TokenType  TokenType
	TokenValue string
}

func (t *TOKEN) New(tokenType TokenType, value string) TOKEN{
	t.TokenType = tokenType
	t.TokenValue = value

	fmt.Printf("New TokenType = %s, TokenValue = %s\n", tokenType, value)

	return *t
}

func (t *TOKEN) toString() {
	fmt.Printf("TOKEN({%s}, {%s})\n", t.TokenType, t.TokenValue);
}
