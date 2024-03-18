package main

import (
	"fmt"
)

func number() {
	var number int
	fmt.Println("Type an integer: ")
	fmt.Scanf("%d", &number)

	if number > 0 {
		fmt.Println("The number is positive.")
		if number == 0 {
			fmt.Println("The number is 0.")
		} else {
			fmt.Println("The number is negative.")
		}
	}
}
