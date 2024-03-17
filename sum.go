package main

import (
	"fmt"
)

func sum() {
	var a = 17
	var b int

	fmt.Println("Enter a number: ")
	fmt.Scan("%d", &b)
	var c = a + b
	fmt.Println("The sum result is: ", c)
}
