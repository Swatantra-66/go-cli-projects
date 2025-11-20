package main

import "fmt"

func main() {
	fmt.Println("Welcome! Do your calculation on goCalculator")

	for {
		fmt.Println(`
1. Addition 
2. Subtraction
3. Multiplication
4. Division
5. Exit`)

		var opt int
		fmt.Print("Enter option: ")
		fmt.Scan(&opt)

		switch opt {
		case 1:
			addition()
		case 2:
			subtraction()
		case 3:
			product()
		case 4:
			division()
		case 5:
			fmt.Println("Thanks for using goCalculator!")
			return
		default:
			fmt.Println("Enter a valid option")
		}
	}
}
