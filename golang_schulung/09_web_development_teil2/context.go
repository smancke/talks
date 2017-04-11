package main

import (
	"fmt"
	"net/http"
	"sync"
)

var context struct {
	data  map[*http.Request]interface{}
	mutex sync.Mutex
}

func main() {
	context.data = make(map[*http.Request]interface{})

	helloWorld := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %v\n", GetUsername(r))
	})
	panic(http.ListenAndServe(":8080", AuthMiddleware(helloWorld)))
}

func AuthMiddleware(delegate http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Username") != "" {
			SetUsername(r, r.Header.Get("Username"))
		} else {
			SetUsername(r, "--not-logged-in--")
		}
		defer ClearUsername(r)
		delegate.ServeHTTP(w, r)
	})
}

func GetUsername(r *http.Request) string {
	context.mutex.Lock()
	defer context.mutex.Unlock()
	return context.data[r].(string)
}

func SetUsername(r *http.Request, username string) {
	context.mutex.Lock()
	defer context.mutex.Unlock()
	context.data[r] = username
}

func ClearUsername(r *http.Request) {
	context.mutex.Lock()
	defer context.mutex.Unlock()
	delete(context.data, r)
}
