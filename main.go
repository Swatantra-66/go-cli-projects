//Number Guessing Game(#1)

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Enter your name: ")
	var userName string
	fmt.Scan(&userName)

	greetUsers(userName)
	difficultyLevel()

	rand.Seed(time.Now().UnixNano())
	guessNum := rand.Intn(100) + 1

	var difficultyChoice uint
	fmt.Println("Enter your choice: ")
	fmt.Scan(&difficultyChoice)

	playGame(difficultyChoice, guessNum) //let's play

}

func greetUsers(userName string) {
	fmt.Printf("\nWelcome %v to the Number Guessing Game!\n", userName)
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You have limited chances to guess the correct number.")
}

func difficultyLevel() {
	fmt.Println("Please select the difficulty level: ")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")
}

func playGame(difficultyChoice uint, guessNum int) {
	var attempts int
	switch difficultyChoice {
	case 1:
		fmt.Println("You have selected the Easy difficulty level.")
		attempts = 10
	case 2:
		fmt.Println("You selected Medium mode (5 attempts).")
		attempts = 5
	case 3:
		fmt.Println("Great! You selected Hard mode (3 attempts).")
		attempts = 3
	default:
		fmt.Println("Invalid choice! Defaulting to Easy mode.")
		attempts = 10
	}

	var num int
	for i := 1; i <= attempts; i++ {
		fmt.Println("Enter your guess: ")
		fmt.Scan(&num)

		if num == guessNum {
			fmt.Printf("Congratulations! You guessed the correct number in %v attempts.\n", i)
			return
		} else if num < guessNum {
			fmt.Printf("Incorrect! The number is greater than %v.\n", num)
		} else {
			fmt.Printf("Incorrect! The number is less than %v.\n", num)
		}
	}
	fmt.Printf("Sorry! The correct number was %d.\n", guessNum)
}
