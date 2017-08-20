package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	healthy := true

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !healthy {
			w.WriteHeader(500)
			fmt.Fprintln(w, `Internal Server Error`)
			return
		}
		fmt.Fprintln(w, `<html><body><h1>Html Page</h1>with some Content</body></html>`)
	})

	http.HandleFunc("/healthy", func(w http.ResponseWriter, r *http.Request) {
		if "POST" == r.Method {
			b, _ := ioutil.ReadAll(r.Body)
			healthy = string(b) == "true"
		}
		fmt.Fprintf(w, "healthy: %v\n", healthy)
	})
	panic(http.ListenAndServe(":80", nil))
}
