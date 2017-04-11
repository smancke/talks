package main

import (
	"fmt"
	"time"
)

func doInBackground(doneChannel chan bool) {
	// do nothing
	doneChannel <- true
}

func main() {
	fmt.Printf("start\n")
	start := time.Now()

	count := 1000000

	doneChannel := make(chan bool, 1)

	for i := 0; i < count; i++ {
		go doInBackground(doneChannel)
	}

	for i := 0; i < count; i++ {
		<-doneChannel
	}
	fmt.Printf("finished %v jobs in %v\n", count, (time.Now().Sub(start)))
}
