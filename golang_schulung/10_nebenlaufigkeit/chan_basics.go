package main

import (
	"fmt"
	"time"
)

func main() {

	//unbuffered()
	//buffered()
	timeAfter()
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

func buffered() {
	ch := make(chan string, 2)

	ch <- "The Answer is "
	ch <- "42"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func timeAfter() {
	timeoutChannel := time.Tick(time.Second)
	for v := range timeoutChannel {
		fmt.Println("One second is elapsed: ", v)
	}
}
