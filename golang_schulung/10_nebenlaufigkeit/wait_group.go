package main

import (
	"fmt"
	"sync"
)

func doInBackground(waitGroup *sync.WaitGroup) {
	fmt.Println("do in backgroud")
	waitGroup.Done()
}

func main() {
	waitGroup := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go doInBackground(waitGroup)
	}

	waitGroup.Wait()

	fmt.Printf("done.")
}
