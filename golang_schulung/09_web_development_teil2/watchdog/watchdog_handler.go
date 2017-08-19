package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

type Watchdog struct {
	watchUrl string
	next     http.Handler
	error    *atomic.Value
}

func NewWatchdog() *Watchdog {
	wd := &Watchdog{}
	wd.error = &atomic.Value{}
	return wd
}

func (wd *Watchdog) setError(val bool) {
	wd.error.Store(val)
}

func (wd *Watchdog) hasError() bool {
	v := wd.error.Load()
	b, _ := v.(bool)
	return b
}

func (wd *Watchdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if wd.hasError {
		fmt.Fprintf(w, "<html><body><h1>We are offline .. please come back later ...</h1></body></html>")
		return
	}
	wd.next.ServeHTTP(w, r)
}

func (wd *Watchdog) watch() {
	failCounter := 0
	for {
		time.Sleep(time.Second)

		resp, err := http.Get(wd.watchUrl)

		if err != nil {
			failCounter++
			fmt.Printf("can't reach server (%v) %q: %v\n", failCounter, wd.watchUrl, err)
			if failCounter > 3 {
				fmt.Printf("serving error page\n")
				wd.hasError() = true
			}
			continue
		}

		if resp.StatusCode != 200 {
			failCounter++
			fmt.Printf("server unhealthy (%v): got code %v\n", failCounter, resp.StatusCode)
			if failCounter > 3 {
				fmt.Printf("serving error page\n")
				wd.hasError() = true
			}
			continue
		}

		if wd.hasError() {
			fmt.Printf("watchdog vaid again\n")
		}
		failCounter = 0
		wd.setError(false)
	}
}
