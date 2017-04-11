package main

import (
	"fmt"
	"testing"
)

func doInBackground(doneChannel chan bool) {
	doneChannel <- true
}

func Benchmark_Creation_Of_Goroutines(b *testing.B) {

	fmt.Printf("testing with %v goroutines\n", b.N)
	doneChannel := make(chan bool)

	for i := 0; i < b.N; i++ {
		go doInBackground(doneChannel)
	}

	for i := 0; i < b.N; i++ {
		<-doneChannel
	}
}
