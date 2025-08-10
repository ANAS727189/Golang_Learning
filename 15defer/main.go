package main

import "fmt"

func main() {
	defer fmt.Println("Goodbye")
	defer fmt.Println("World")
	fmt.Println("Hello")
	defer fmt.Println("from")
	defer fmt.Println("Go")
}
