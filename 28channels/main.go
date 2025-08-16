package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")

	// mych := make(chan int, 2) Normal channel
	mych := make(chan int, 2) // Buffered channel
	wg := &sync.WaitGroup{}

	// mych <- 5
	// fmt.Println("Value sent to channel:", <-mych)

	wg.Add(2)

	go func(ch chan<- int, wg *sync.WaitGroup) { //Send only
		defer wg.Done()
		mych <- 5
		mych <- 10
		mych <- 15
		close(ch) // Closing the channel after sending values
	}(mych, wg)

	go func(ch <-chan int, wg *sync.WaitGroup) { //Receive only
		defer wg.Done()
		fmt.Println("Value received from channel:", <-ch)
		fmt.Println("Value received from channel:", <-ch)
		val, isChannelOpen := <-ch
		if isChannelOpen {
			fmt.Println("Status: ", isChannelOpen)
			fmt.Println("Value received from channel:", val)
		} else {
			fmt.Println("Status: ", isChannelOpen)
			fmt.Println("Channel is closed!")
		}
	}(mych, wg)

	wg.Wait()
}
