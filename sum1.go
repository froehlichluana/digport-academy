package main

import (
	"fmt"
)

func sumtest() {
	var a = 17
	var b int
	var sum1 = a + b

	fmt.Println("Enter a number: ")
	fmt.Scanf("%d", &b)
	fmt.Println("The result is ", sum1)
}
