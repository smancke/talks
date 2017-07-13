package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: reverse_proxy targetUrl watchUrl")
		return
	}
	targetUrl, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Printf("invalid url %q", os.Args[1])
		return
	}

	watchUrl := os.Args[2]

	watchdog := &Watchdog{
		watchUrl: watchUrl,
		next: &LocationFix{
			createReverseProxy(targetUrl),
		},
	}
	go watchdog.watch()

	panic(http.ListenAndServe(":8080", watchdog))
}

func createReverseProxy(target *url.URL) http.Handler {
	proxy := httputil.NewSingleHostReverseProxy(target)
	director := proxy.Director
	proxy.Director = func(req *http.Request) {
		req.Host = target.Host
		director(req)
	}
	return proxy
}
