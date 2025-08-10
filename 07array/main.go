package main

import "fmt"

func main() {
	fmt.Println("Hello, Today we will learn about arrays in Go!")

	var arr [5]int
	arr[0] = 10
	arr[1] = 20
	// arr[2] = 30
	arr[3] = 40
	arr[4] = 50

	fmt.Println("Array elements are:", arr)
	fmt.Println("Length of the array is:", len(arr))
	fmt.Println("Capacity of the array is:", cap(arr))

	var fruitList = [4]string{"Apple", "Banana", "Cherry", "Date"}
	fmt.Println("Fruit list is:", fruitList)
}
