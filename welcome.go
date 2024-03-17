package main

import "fmt"

func welcome() {
	var name string

	fmt.Println("What's your name? ")
	fmt.Scan(&name)
	fmt.Println("Welcome, " + name)
	fmt.Printf("Welcome, %s !/n", name)

}
