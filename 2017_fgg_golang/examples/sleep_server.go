package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

func main() {
	var totalCalls int32 = 0

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Millisecond)
		callNumber := atomic.AddInt32(&totalCalls, 1)
		fmt.Fprintf(w, "call %v\n", callNumber)
	})

	panic(http.ListenAndServe(":8080", handler))
}
