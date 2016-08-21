package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type Queue chan string

func main() {

	callCount := 30000
	workerCount := 500
	http.DefaultClient.Transport = &http.Transport{
		MaxIdleConnsPerHost: 500,
	}

	fmt.Println("start")
	start := time.Now()

	waitgroup := &sync.WaitGroup{}

	queue := make(Queue)

	for workerCount > 0 {
		go workerLoop(queue, waitgroup)
		waitgroup.Add(1)
		workerCount--
	}

	for i := 0; i < callCount; i++ {
		queue <- "http://127.0.0.1:8080"
	}
	close(queue)

	waitgroup.Wait()
	fmt.Printf("%v calls took %v\n", callCount, time.Since(start))
}

func workerLoop(queue Queue, waitgroup *sync.WaitGroup) {
	for url := range queue {
		if resp, err := http.Get(url); err != nil {
			log.Println(err.Error())
		} else {
			ioutil.ReadAll(resp.Body)
			resp.Body.Close()
		}
	}
	waitgroup.Done()
}
