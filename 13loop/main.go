package main

import "fmt"

func main() {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println("Days of the week:", days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }
	for d := range days {
		fmt.Println(days[d])
	}

	for _, day := range days {
		fmt.Printf("Index is _  and value is: %v\n", day)
	}

	rogueValue := 1
	for rogueValue < 10 {
		if rogueValue == 5 {
			rogueValue++
			continue
		}
		fmt.Println("Value is: ", rogueValue)
		rogueValue++
	}
}
