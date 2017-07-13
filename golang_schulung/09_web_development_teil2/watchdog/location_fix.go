package main

import (
	"fmt"
	"net/http"
	"net/url"
)

type LocationFix struct {
	next http.Handler
}

func (l *LocationFix) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	l.next.ServeHTTP(&LocationFixResponseWriter{w, false}, r)
}

type LocationFixResponseWriter struct {
	http.ResponseWriter
	headerSend bool
}

func (w *LocationFixResponseWriter) WriteHeader(statusCode int) {
	fixLocation(w.ResponseWriter)
	w.ResponseWriter.WriteHeader(statusCode)
	w.headerSend = true
}

func fixLocation(resp http.ResponseWriter) {
	l := resp.Header().Get("Location")
	if l != "" {
		if u, err := url.Parse(l); err == nil {
			p := u.Path
			if u.RawQuery != "" {
				p += "?" + u.RawQuery
			}
			if u.Fragment != "" {
				p += "#" + u.Fragment
			}
			resp.Header().Set("Location", p)
		} else {
			fmt.Println(err)
		}
	}
}

func (w *LocationFixResponseWriter) Write(bytes []byte) (int, error) {
	if !w.headerSend {
		w.WriteHeader(200)
	}
	return w.ResponseWriter.Write(bytes)
}
