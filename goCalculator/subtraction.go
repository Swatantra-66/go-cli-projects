package main

import (
	"fmt"
	"log"
)

func subtraction() {
	var num1 int64
	var num2 int64
	fmt.Printf("Enter num 1: ")
	_, err := fmt.Scan(&num1)
	if err != nil {
		log.Fatalf("Invalid input: %v", err)
	}

	fmt.Printf("Enter num 2: ")
	_, err1 := fmt.Scan(&num2)
	if err1 != nil {
		log.Fatalf("Invalid input: %v", err1)
	}

	result := num1 - num2
	fmt.Printf("%v\n", result)
}
