package main

import (
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"sort"
	"sync"
	"syscall"
)

var aList = make([]string, 0)
var aMutex = new(sync.Mutex)

func main() {
	runtime.SetBlockProfileRate(50)

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	for i := 0; i < 50; i++ {
		go doSomething()
	}

	sigc := make(chan os.Signal)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("Got singal '%v' .. exit now", <-sigc)
}

func doSomething() {
	for {
		aMutex.Lock()
		for i := 0; i < 10000; i++ {
			aList = append(aList, randSeq(50))
		}
		sort.Strings(aList)
		if len(aList) > 10000000 {
			aList = aList[0:10000000]
		}
		log.Printf("list len %v", len(aList))
		aMutex.Unlock()
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
