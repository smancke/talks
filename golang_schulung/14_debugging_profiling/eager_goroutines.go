package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	for i := 0; i < 500; i++ {
		go saySomething()
	}

	sigc := make(chan os.Signal)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("Got singal '%v' .. exit now\n", <-sigc)

}

func saySomething() {
	for {
		//fmt.Println("i want to say something")
		time.Sleep(time.Millisecond)
	}
}
