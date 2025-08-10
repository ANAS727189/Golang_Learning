package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	postReq()
	fmt.Println("Post request completed")
}

func postReq() {
	const myUrl = "http://localhost:8000/post"

	reqBody := strings.NewReader(`
	{
	"courseName": "Golang",
	"price": 0,
	"platform": "Coursera"}
	`)
	res, err := http.Post(myUrl, "application/json", reqBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)
	fmt.Println("Response from server:", string(content))
	fmt.Println("Status Code:", res.StatusCode)
}
