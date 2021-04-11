package main

import (
	"bufio"
	"fmt"
	"github.com/santoshbachar/go-simple-interpreter/interpreter"
	"os"
)

func main() {
	fmt.Println("Hello, world")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	fmt.Print("text = ", text)

	var interpreter interpreter.Interpreter

	interpreter.Init(text)

	result := interpreter.Expr()

	fmt.Println("interpreter = ", interpreter)
	fmt.Println("result = ", result)
}
