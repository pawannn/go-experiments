package main

import "net/http"

type Server struct {
	c *Cache
}

func NewServer(c *Cache) *Server {
	return &Server{
		c: c,
	}
}

func main() {
	c := NewTTLCache()
	server := NewServer(c)

	http.HandleFunc("/cache/", server.HandleRequest)
	http.ListenAndServe(":8080", nil)
}
