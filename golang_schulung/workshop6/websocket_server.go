package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var webSocketUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := webSocketUpgrader.Upgrade(w, r, nil)
		if err != nil {
			panic(err)
		}

		for {
			_, bytes, err := c.ReadMessage()
			if err != nil {
				return
			}
			c.WriteMessage(websocket.BinaryMessage, bytes)
		}
	})

	panic(http.ListenAndServe(":8080", handler))
}
