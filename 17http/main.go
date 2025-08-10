package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("HTTP requests in Go")

	url := "http://youtube.com"

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	// fmt.Println("Response status:", res)
	defer res.Body.Close()

	databytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println("Content :", content)
}
