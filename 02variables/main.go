package main

import (
	"fmt"
)

const LoginStudent string = "logged in" // public
// (here capital L is necessary)

func main() {
	var username string = "Anas"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)

	var isVerified bool = true
	fmt.Println(isVerified)
	fmt.Printf("Variable is of type: %T\n", isVerified)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T\n", smallValue)

	var largeDecimal float64 = 43.55422445865
	fmt.Println(largeDecimal)
	fmt.Printf("Variable is of type: %T\n", largeDecimal)

	var name = "Anas"
	fmt.Println(name)

	numberOfStudents := 10000
	// := is called walrus operator (it can only be used inside a method/function/main)
	fmt.Println(numberOfStudents)

	fmt.Println(LoginStudent)

}
