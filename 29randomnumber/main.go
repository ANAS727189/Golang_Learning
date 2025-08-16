package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	// "time"
	"crypto/rand"
)

func main() {
	fmt.Println("Random number in Golang")

	//RANDOM NUMBER USING MATH PACKAGE
	// rand.Seed(time.Now().UnixNano())
	// mathRandomNumber := rand.Intn(6) + 1 // random number from 1 to 6
	// fmt.Println("Random number:", mathRandomNumber)

	//RANDOM NUMBER USING CRYPTO
	cryptoRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(7)) // random number from 0 to 6
	fmt.Println("Random number:", cryptoRandomNumber)

}
