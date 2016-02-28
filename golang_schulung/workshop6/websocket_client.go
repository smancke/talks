package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"os"
	"strings"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8080", nil)
	if err != nil {
		panic(err)
	}
	conn.WriteMessage(websocket.TextMessage, []byte(strings.Join(os.Args[1:], " ")))
	_, body, err := conn.ReadMessage()
	fmt.Println(string(body))
}
