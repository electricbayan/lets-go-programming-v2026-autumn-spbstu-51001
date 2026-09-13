package main

import (
	"fmt"
)

func main() {
	var firstArg int
	_, err := fmt.Scan(&firstArg)
	if err != nil {
		fmt.Println("Invalid first operand")
	}

	var secondArg int
	_, err = fmt.Scan(&secondArg)
	
	if err != nil {
		fmt.Println("Invalid second operand")
	}

	var operator string
	_, err = fmt.Scan(&operator)
	
	switch operator {
	case "+":
		fmt.Println(firstArg + secondArg)
	case "-":
		fmt.Println(firstArg - secondArg)
	case "*":
		fmt.Println(firstArg * secondArg)
	case "/":
		if (secondArg != 0) {
			fmt.Println(firstArg / secondArg)
		} else {
			fmt.Println("Division by zero")
		}
	default:
		fmt.Println("Invalid operation")
	}
}