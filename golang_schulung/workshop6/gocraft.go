package main

import (
	"fmt"
	"github.com/gocraft/web"
	"net/http"
)

type Context struct {
	Username string
}

func (c *Context) Auth(w web.ResponseWriter, r *web.Request, next web.NextMiddlewareFunc) {
	if r.Header.Get("Username") != "" {
		c.Username = r.Header.Get("Username")
	} else {
		c.Username = "--not-logged-in--"
	}
	next(w, r)
}

func (c *Context) Hello(w web.ResponseWriter, r *web.Request) {

	fmt.Fprintf(w, "Hello %v\n", c.Username)
}

func main() {
	mux := web.New(Context{}).
		Middleware(web.LoggerMiddleware).
		Middleware((*Context).Auth).
		Get("/", (*Context).Hello)

	panic(http.ListenAndServe(":8080", mux))
}
