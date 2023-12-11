// main.go
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	WriteBufferSize: 1024,
	ReadBufferSize:  1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer connection.Close()

	for {
		messageType, p, err := connection.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		if p != nil {
			fmt.Println(p)
			fmt.Println("Request received !")
		}

		if err := connection.WriteMessage(messageType, p); err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)

	fmt.Println("WebSocket server running on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}

var upgrader2 = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func handleRequest2(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader2.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		switch string(p) {

		case "create":
			err = conn.WriteMessage(messageType, []byte("Create request received"))
		case "update":
			err = conn.WriteMessage(messageType, []byte("Update request received"))
		case "delete":
			err = conn.WriteMessage(messageType, []byte("Delete request received"))
		default:
			err = conn.WriteMessage(messageType, []byte("Unknown request"))
		}

		if err != nil {
			fmt.Println(err)
			return
		}
	}
}