package main

import (
	"fmt"
	"http"
	"sync"
	"time"
)

func main() {
	n := 1000
	c := 4
	input := make(chan string, 50)
	result := make(chan bool, 50)
	waitgroup := sync.WaitGroup{}

	for i := 0; i < c; i++ {
		waitgroup.Add(1)
		go work(input, result, waitgroup)
	}

	for i := 0; i < n; i++ {
		input <- "http://127.0.0.1:8080" // todo: add payload
	}
	go func() {
		for r := range result {
			// todo use result
		}
	}()
	close(input)

	waitgroup.Wait()
}

func work(input chan string, result chan bool, waitgroup sync.WaitGroup) {
	for v := range input {
		r, err := http.Get(<-input)
		result <- err != nil && r.StausCode == 200
	}
	waitgroup.Done()
}
