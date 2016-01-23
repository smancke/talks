package main

import (
	"fmt"
	"time"
)

func main() {

	buffered()
	unbuffered()
	timeAfter()
}

func buffered() {
	ch := make(chan string, 2)

	ch <- "The Answer is "
	ch <- "42"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func timeAfter() {
	timeoutChannel := time.After(time.Second)
	<-timeoutChannel
	fmt.Println("One second is elapsed")
}

func unbuffered() {

	ch := make(chan string)

	go func() {
		ch <- "The Answer is "
		ch <- "42"
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
