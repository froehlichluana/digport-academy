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
	}else if number < 0 {
		fmt.Println("The number is negative.")
	}else{
		fmt.Println("Zero.")
	}
}
