package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func main() {

	mux := httprouter.New()
	mux.GET("/hello/:name", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Fprintf(w, "Hello %v\n", params.ByName("name"))
	})

	mux.GET("/matchall/*path", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Fprintf(w, "The path %v\n", params.ByName("path"))
	})

	mux.POST("/echo", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		body, _ := ioutil.ReadAll(r.Body)
		w.Write(body)
	})

	mux.HandlerFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from some usual handler\n")
	})

	mux.ServeFiles("/static/*filepath", http.Dir("."))

	mux.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Ups ... no such page\n")
	})

	mux.PanicHandler = func(w http.ResponseWriter, r *http.Request, error interface{}) {
		fmt.Fprintf(w, "Ups, there was an internal error: %v\n", error)
	}

	mux.GET("/panic", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		panic("PAAAANNIIICCC")
	})

	panic(http.ListenAndServe(":8080", mux))
}
