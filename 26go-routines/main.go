package main

import (
	"fmt"
	"net/http"
	"sync"
	// "time"
)

var wg sync.WaitGroup
var mut sync.Mutex // Mutex to protect shared resources

var signals []string //Example slice to demonstrate mutex usage

func main() {
	wg.Add(2)
	go greeter("Hello") // Use "go" keyword for Go routine
	greeter("World")
	websiteList := []string{
		"https://lco.dev",
		"http://google.com",
		"https://go.dev",
		"https://fb.com",
		"https://github.com",
	}

	for _, website := range websiteList {
		go getStatusCode(website)
		wg.Add(1)
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println(signals)
	fmt.Println("All goroutines finished executing.")
}

func greeter(s string) {
	defer wg.Done() // Ensure that the wait group counter is decremented
	for i := 0; i < 6; i++ {
		// time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	mut.Lock()
	signals = append(signals, endpoint) // Example of shared resource usage
	mut.Unlock()
	fmt.Printf("%d Status Code for %s:\n", res.StatusCode, endpoint)
}
