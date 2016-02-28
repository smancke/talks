package main

import (
	"fmt"
	"github.com/gorilla/context"
	"net/http"
)

var Username = "Username"

func main() {

	helloWorld := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %v\n", context.Get(r, Username))
	})
	panic(http.ListenAndServe(":8080", context.ClearHandler(AuthMiddleware(helloWorld))))
}

func AuthMiddleware(delegate http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Username") != "" {
			context.Set(r, Username, r.Header.Get("Username"))
		} else {
			context.Set(r, Username, "--not-logged-in--")
		}
		delegate.ServeHTTP(w, r)
	})
}
