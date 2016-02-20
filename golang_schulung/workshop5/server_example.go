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

	server := http.Server{
		Addr:    ":8080",
		Handler: handler("Hello World"),
	}

	server.ListenAndServe()
}
