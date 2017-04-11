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
	panic(http.ListenAndServe(":8080", handler("Hallo Welt")))
}
