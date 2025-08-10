package main

func main() {
	result := adder(1, 2, 3, 4, 5)
	println("The sum is:", result) // Output: The sum is: 15
}

func adder(val ...int) int {
	ans := 0
	for _, v := range val {
		ans += v
	}
	return ans
}
