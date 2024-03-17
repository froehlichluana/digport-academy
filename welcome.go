package main

import (
	"fmt"
)

func Welcome() {
	var name string

	fmt.Println("What's your name? ")
	fmt.Scan(&name)
	fmt.Println("Welcome, " + name)

}
