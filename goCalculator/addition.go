package main

import (
	"fmt"
	"log"
)

func addition() {
	var n int
	fmt.Print("How many numbers do you want to add? ")
	_, err := fmt.Scan(&n)
	if err != nil || n <= 0 {
		log.Fatalf("Invalid count: %v", err)
	}

	var sum int64
	for i := 1; i <= n; i++ {
		var x int64
		fmt.Printf("Enter number %d: ", i)
		_, err := fmt.Scan(&x)
		if err != nil {
			log.Fatalf("Invalid number at item %d: %v", i, err)
		}
		sum += x
	}

	fmt.Printf("Sum is: %d\n", sum)
}
