package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("File operations in Go")
	content := "This is a sample content for the file."

	file, err := os.Create("./my-first-go-file.txt")
	checkError(err)

	length, err := io.WriteString(file, content)
	checkError(err)

	fmt.Printf("Length of written content: %d bytes\n", length)
	defer file.Close()
	readFile("./my-first-go-file.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkError(err)
	// fmt.Println("File content: \n", databyte)
	fmt.Println("File content: \n", string(databyte))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
