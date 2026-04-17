package main

import (
	"fmt"
	"log"
	"net/http"
)

// FOR DEMO ONLY
func StartServer(addr string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Response from backend %s\n", addr)
	})

	log.Println("Starting backend on", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
