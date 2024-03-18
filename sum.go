package main

import (
	"fmt"
)

func sum1() {
	var a = 17
	var b int

	fmt.Println("Enter a number: ")
	fmt.Scanf("%d", &b)
	var c = a + b
	fmt.Println("The sum result is: ", c)
}
