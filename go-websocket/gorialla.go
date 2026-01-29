package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func NewHub() *Hub {
	return &Hub{
		client: make(map[*websocket.Conn]struct{}),
	}
}

func (h *Hub) HandlerWs(w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer conn.Close()

	h.Register(conn)
}

func (h *Hub) ReadLoop(conn *websocket.Conn) {
	defer func() {
		h.Unregister(conn)
		conn.Close()
	}()

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		h.BroadCast(msgType, msg)
	}
}

func Listen() error {
	hub := NewHub()
	http.HandleFunc("/ws", hub.HandlerWs)
	fmt.Println("WS listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return err
	}

	return nil
}
