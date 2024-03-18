package main

import (
	"fmt"
)

func main() {
	var a = 17
	var b int
	var c = a + b

	fmt.Println("Enter a number: ")
	fmt.Scanf("%d", &b)
	fmt.Println("The result is ", &c)
}
