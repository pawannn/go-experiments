package main

import (
	"log"
	"net/http"
)

func main() {
	backendPorts := []string{":9091", ":9092", ":9093"}

	var backends []*Backend
	for _, port := range backendPorts {
		addr := "http://localhost" + port
		backend, err := NewBackend(addr)
		if err != nil {
			continue
		}

		backends = append(backends, backend)

		go StartServer(port)
	}

	lb := NewLoadBalancer(backends)

	log.Println("Load balancer running on :8080")
	http.ListenAndServe(":8080", lb)
}
