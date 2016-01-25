package main

import (
	"fmt"
	"time"
)

func main() {

	sendWithTimeout()
	readNonBlocking()
}

func sendWithTimeout() {
	ch := make(chan string)

	select {
	case ch <- "a message":
	case <-time.After(time.Millisecond):
		fmt.Println("timeout on writing to channel.")
		return
	}
}

func readNonBlocking() {
	ch := make(chan string, 2)

	ch <- "The Answer is "
	ch <- "42"

	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		default:
			fmt.Println("no input available.")
			return
		}
	}
}
