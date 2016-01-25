package main

import (
	"fmt"
)

func main() {

	byRange()
	byClose()
}

func byClose() {
	ch := make(chan string)

	go func() {
		ch <- "The Answer is "
		ch <- "42"
		close(ch)
	}()

	for {
		msg, channelOpen := <-ch
		if !channelOpen {
			break
		}
		fmt.Println(msg)
	}
}

func byRange() {
	ch := make(chan string, 2)

	ch <- "The Answer is "
	ch <- "42"
	close(ch)

	for v := range ch {
		fmt.Printf("v=%v\n", v)
	}
}
