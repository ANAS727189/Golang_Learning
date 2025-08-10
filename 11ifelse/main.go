package main

import "fmt"

func main() {
	fmt.Println("Welcome to the if-else example in Go!")

	loginCnt := 20
	var res string

	if loginCnt < 10 {
		res = "Regular User"
	} else if loginCnt >= 10 && loginCnt < 20 {
		res = "Moderate User"
	} else {
		res = "Power User"
	}
	fmt.Println("User Level:", res)
}
