package main

import (
	"math/rand"
	"sync"
	"time"
)

var aMap = make(map[string]string, 0)
var aMutex = new(sync.Mutex)

func main() {
	for i := 0; i < 2; i++ {
		go modifyMap()
	}

	time.Sleep(time.Second)
}

func modifyMap() {
	for i := 0; i < 100; i++ {
		aMutex.Lock()
		key := randSeq(10)
		value := randSeq(10)
		aMap[key] = value
		delete(aMap, key)
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
