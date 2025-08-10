package main

import "fmt"

func main() {
	fmt.Println("Today we will learn about pointers.")
	var ptr *int
	fmt.Println("Default value of pointer is: ", ptr)
	myName := "Anas"
	var check_ptr *string = &myName
	fmt.Println("My name using pointers is: ", check_ptr)

	myNum := 42
	var new_ptr = &myNum
	fmt.Println("Address of myNum is: ", new_ptr)
	fmt.Println("Value of myNum is: ", *new_ptr)

	*new_ptr += 1
	fmt.Println("New value is: ", myNum)
}
