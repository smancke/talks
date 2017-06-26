package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const gracePeriod = time.Second * 5

func main() {
	handlerChain := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 5)
		fmt.Fprintf(w, "Hello Welt\n")
	})
	httpSrv := &http.Server{Addr: ":8080", Handler: handlerChain}

	go func() {
		if err := httpSrv.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				fmt.Println("not accepting new connections")
			} else {
				fmt.Printf("error %v", err)
				os.Exit(1)
			}
		}
	}()

	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	sig := <-stop
	fmt.Printf("got %v, shutdown gracefully, now\n", sig)

	ctx, ctxCancel := context.WithTimeout(context.Background(), gracePeriod)

	httpSrv.Shutdown(ctx)
	fmt.Println("down")
	ctxCancel() // not needed, but good style
}
