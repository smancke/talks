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
	fmt.Printf("start")
	start := time.Now()

	count := 1000000

	doneChannel := make(chan bool)

	for i := 0; i < count; i++ {
		go doInBackground(doneChannel)
	}

	for i := 0; i < count; i++ {
		<-doneChannel
	}
	fmt.Printf("finished %v jobs in %v", count, (time.Now().Sub(start)))
}
