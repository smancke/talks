package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	chain := LoggingMiddleware(AccessMiddleware(helloWorld))
	panic(http.ListenAndServe(":8080", chain))
}

var helloWorld = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World with %v\n", r.Proto)
})

func AccessMiddleware(delegate http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/restricted") {
			http.Error(w, "Access Denied", 409)
		} else {
			delegate.ServeHTTP(w, r)
		}
	})
}

func LoggingMiddleware(delegate http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v %v", r.RemoteAddr, r.Method, r.URL.Path)
		delegate.ServeHTTP(w, r)
	})
}
