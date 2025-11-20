package main

import "fmt"

func division() {
	var num1 int
	var num2 int
	fmt.Printf("Enter num 1: ")
	fmt.Scan(&num1)

	fmt.Printf("Enter num 2: ")
	fmt.Scan(&num2)

	if num2 == 0 {
		fmt.Println("Cannot divide by zero")
	} else {
		result := num1 / num2
		fmt.Printf("%v\n", result)
	}
}
