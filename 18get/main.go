package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	getReq()
}

func getReq() {
	const myUrl = "https://youtube.com"

	res, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println("Status Code:", res.StatusCode)

	content, _ := io.ReadAll(res.Body)
	// fmt.Println("Response Body:", string(content))
	fmt.Println("Response Length:", len(content))
}
