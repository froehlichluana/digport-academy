package main

import (
	"fmt"
)

func countdown() {

	for contador := 10; contador > 0; contador --{
		fmt.Println("Count down", contador)
	}
	fmt.Println("Happy New Year!")
}