package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World with %v\n", r.Proto)
	})
	panic(http.ListenAndServeTLS(":8443", "server.pem", "server.key", nil))
}
