package main

import (
	"fmt"
	"time"
)

func startWorker(name string, startSignal chan bool) {
	<-startSignal
	fmt.Printf("Worker %v got start singal\n", name)
}

func main() {

	ch := make(chan bool)

	go startWorker("Worker 1", ch)
	go startWorker("Worker 2", ch)
	go startWorker("Worker 3", ch)

	close(ch)

	time.Sleep(time.Millisecond)
}
