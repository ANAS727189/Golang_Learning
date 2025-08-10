package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to learning go course with me.")
	fmt.Println("Please first tell me number between 1-5 defining your level of understanding Golang")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Your input is: " + input)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("Error parsing input, please enter a valid number between 1-5", err)
		panic(err)
	}
	fmt.Println("Thanks for your rating, but we have decreased your understanding by 1, your new rating is: ", numRating-1)

}
