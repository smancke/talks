package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {

	urlMainPage, _ := url.Parse("http://shop-article-list/")
	http.Handle("/", httputil.NewSingleHostReverseProxy(urlMainPage))

	urlOfTheBasked, _ := url.Parse("http://cart/")
	http.Handle("/cart", httputil.NewSingleHostReverseProxy(urlOfTheBasked))

	panic(http.ListenAndServe(":80", nil))
}
