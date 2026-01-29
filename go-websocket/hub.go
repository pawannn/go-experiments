package main

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	mu     sync.Mutex
	client map[*websocket.Conn]struct{}
}

func NewHubManager() *Hub {
	return &Hub{
		client: make(map[*websocket.Conn]struct{}),
	}
}

func (h *Hub) Register(conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.client[conn] = struct{}{}

	h.ReadLoop(conn)
}

func (h *Hub) Unregister(conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.client, conn)
}

func (h *Hub) BroadCast(messageType int, msg []byte) {
	for conn := range h.client {
		if err := conn.WriteMessage(messageType, msg); err != nil {
			fmt.Println(err)
			go conn.Close()
			go h.Unregister(conn)
		}
	}
}
