package main

import (
	"fmt"
)

func main() {
	var b int
	a := 1

	fmt.Printf("Enter a number")
	fmt.Scanf("%d", &b)
	sum := a + b

	fmt.Println("My return:", sum)

}
