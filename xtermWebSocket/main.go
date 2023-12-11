package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

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
func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "\n=====\n")
	fmt.Fprint(w, "0k")
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/server", handleRequest)
	fmt.Println("Server started at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
