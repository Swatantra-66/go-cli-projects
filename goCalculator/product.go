package main

import (
	"fmt"
	"log"
)

func product() {
	var n int
	fmt.Print("How many numbers do you want to multiply? ")
	_, err := fmt.Scan(&n)
	if err != nil || n <= 0 {
		log.Fatalf("Invalid count: %v", err)
	}

	var result int64
	result = 1
	for i := 1; i <= n; i++ {
		var x int64
		fmt.Printf("Enter number %d: ", i)
		_, err := fmt.Scan(&x)
		if err != nil {
			log.Fatalf("Invalid number at item %d: %v", i, err)
		}
		result *= x
	}
	fmt.Printf("Product is: %d\n", result)
}
