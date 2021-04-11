package interpreter

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/santoshbachar/go-simple-interpreter/interpreter/lexer"
)

// Interpreter ...
type Interpreter struct {
	text         string
	pos          int
	currentToken lexer.TOKEN
}

func (i *Interpreter) Init(text string) Interpreter {
	i.text = text
	i.pos = 0
	i.currentToken = lexer.TOKEN{}

	return *i
}

func (i *Interpreter) error() error {
	fmt.Println("Error: Interpreter")
	return errors.New("hello")
}

func (i *Interpreter) getNextToken() (error, lexer.TOKEN) {
	var text = i.text
	var token lexer.TOKEN

	if i.pos > len(text)-1 {
		return nil, token.New(lexer.EOF, "")
	}

	var currentChar = text[i.pos]

	if lexer.IsCharADigit(currentChar) {
		token.New(lexer.INTEGER, string(currentChar))
		i.pos += 1
		return nil, token
	}

	if currentChar == '+' {
		token.New(lexer.SUM, string (currentChar))
		i.pos += 1
		return nil, token
	}

	return i.error(), lexer.TOKEN{}
}

func (i *Interpreter) eat(tokenType lexer.TokenType) {

	fmt.Printf("currentToken.TokenType = %s, TokenType = %s\n", i.currentToken.TokenType, tokenType)

	if i.currentToken.TokenType == tokenType {
		_, i.currentToken = i.getNextToken()
	} else {
		i.error()
	}
}

func (i *Interpreter) Expr() int {
	var err error
	err ,i.currentToken = i.getNextToken()

	if err != nil {

	}

	left, _ := strconv.Atoi(i.currentToken.TokenValue)
	i.eat(lexer.INTEGER)
	fmt.Println("left = ", left)

	op := i.currentToken
	i.eat(lexer.SUM)
	fmt.Println("op in Expr() ", op)

	right, _ := strconv.Atoi(i.currentToken.TokenValue)
	i.eat(lexer.INTEGER)
	fmt.Println("right = ", right)

	result := left + right

	return result
}
