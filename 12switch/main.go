package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the switch example in Go!")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of diceNumber:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You rolled a one! Try again.")
	case 2:
		fmt.Println("You rolled a two! Not bad.")
	case 3:
		fmt.Println("You rolled a three! Good job.")
	case 4:
		fmt.Println("You rolled a four! Nice roll.")
	case 5:
		fmt.Println("You rolled a five! Almost there.")
	case 6:
		fmt.Println("You rolled a six! Excellent!")
	default:
		fmt.Println("Invalid roll. Please try again.")
	}

}
