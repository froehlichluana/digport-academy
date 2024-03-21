package main

import (
	"fmt"
)

func main() {

	var a int
	var b int
	var operator string

	fmt.Println("Enter number a")
	fmt.Scan(&a)

	fmt.Println("Enter number b")
	fmt.Scan(&b)

	fmt.Println("Enter operator: +, -, *, /")
	fmt.Scan(&operator)

	switch operator{
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		fmt.Println(a / b)
	default:
		fmt.Println("Invalid operator!")

	}

}