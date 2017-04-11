package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.Handle("/world", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World\n")
	}))
	mux.Handle("/bonn", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Bonn\n")
	}))

	http.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mux.ServeHTTP(w, r)
		fmt.Fprintf(w, "Ich werde immer aufgerufen\n")
	}))

	err := http.ListenAndServe(":8080", http.DefaultServeMux)
	if err != nil {
		panic(err)
	}
}
