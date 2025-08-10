package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	performPost()
	fmt.Println("Post request completed successfully!")
}

func performPost() {
	const myUrl = "http://localhost:8000/post-form"

	data := url.Values{}
	data.Add("name", "Anas Khan")
	data.Add("age", "25")
	data.Add("city", "Kanpur")
	res, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)
	fmt.Println("Response from server:", string(content))
	fmt.Println("Status Code:", res.StatusCode)
}
