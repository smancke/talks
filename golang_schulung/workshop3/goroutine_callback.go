package main

import (
	"fmt"
)

type request struct {
	msg      string
	callback chan string
}

func echoRoutine(requestChannel chan request) {
	for {
		request := <-requestChannel
		request.callback <- request.msg
	}
}

func main() {

	requestChannel := make(chan request, 1)
	go echoRoutine(requestChannel)

	response := make(chan string)

	requestChannel <- request{"Hello ", response}
	fmt.Print(<-response)

	requestChannel <- request{"World", response}
	fmt.Print(<-response)
}
