package main

import (
	"fmt"
	"os"
	"runtime"
)

func doInBackground() {
	for i := 0; i < 100; i++ {
		fmt.Println("in background")
		//runtime.Gosched()
	}
}

func main() {

	runtime.GOMAXPROCS(4)

	go doInBackground()

	for i := 0; i < 100; i++ {
		fmt.Println("in foreground")
		//runtime.Gosched()
	}

	os.Exit(0)
}
