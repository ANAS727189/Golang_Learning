package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Mango", "Guava", "Banana", "Peach"}
	fmt.Println("Fruit list is:", fruitList)
	fruitList = append(fruitList, "Tomato", "Orange")
	fmt.Println("New Fruit list is:", fruitList)

	fruitList = fruitList[1:]
	fmt.Println("After removing first element, Fruit list is:", fruitList)

	fruitList = fruitList[1:3]
	//it runs through from index 1 to index 3 (exclusive)
	fmt.Println("After slicing, Fruit list is:", fruitList)

	highScores := make([]int, 5)
	highScores[0] = 154
	highScores[1] = 120
	highScores[2] = 30
	highScores[3] = 10
	highScores[4] = 50
	fmt.Println("High scores are:", highScores)
	fmt.Println("Length of high scores is:", len(highScores))
	highScores = append(highScores, 600, 700)
	fmt.Println("New high scores are:", highScores)
	fmt.Println("Length of new high scores is:", len(highScores))
	sort.Ints(highScores)
	fmt.Println("Sorted high scores are:", highScores)

	var courses = []string{"Go", "Python", "Java", "C++", "JavaScript"}
	fmt.Println("Courses are:", courses)
	var idx int = 2
	courses = append(courses[:idx], courses[idx+1:]...)
	fmt.Println("After removing course at index", idx, "Courses are:", courses)
}
