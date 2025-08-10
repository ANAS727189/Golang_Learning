package main

import "fmt"

//No inheritance in golang, no super/parent etc

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	Anas := User{
		Name:   "Anas",
		Email:  "anas@example.com",
		Status: true,
		Age:    30,
	}

	fmt.Println("User Details:", Anas)
	fmt.Printf("Anas details are: %+v\n", Anas)
	fmt.Printf("Name is: %v & Email is: %v\n", Anas.Name, Anas.Email)
}
