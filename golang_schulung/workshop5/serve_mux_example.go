package main

import (
	"fmt"
	"net/http"
)

type handler string

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v\n", h)
}

func main() {

	mux := http.NewServeMux()
	mux.Handle("/hello", handler("Hallo Welt"))
	mux.HandleFunc("/answer", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "42\n")
	})

	panic(http.ListenAndServe(":8080", mux))
}
