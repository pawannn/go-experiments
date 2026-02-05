package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"sync"
)

type LoadBalancer struct {
	backends []*Backend
	current  int
	mu       sync.Mutex
}

func NewLoadBalancer(backends []*Backend) *LoadBalancer {
	return &LoadBalancer{
		backends: backends,
		current:  0,
	}
}

func (lb *LoadBalancer) NextBackendServer() *Backend {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	n := len(lb.backends)
	for range n {
		lb.current = (lb.current + 1) % n
		if lb.backends[lb.current].IsAlive() {
			return lb.backends[lb.current]
		}
	}

	return nil
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	backend := lb.NextBackendServer()

	if backend == nil {
		http.Error(w, "service unavailable", http.StatusServiceUnavailable)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(backend.url)
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		fmt.Println("Backend error : ", backend.url)
		backend.SetAlive(false)
		http.Error(w, "service unavailable", http.StatusBadGateway)
	}

	proxy.ServeHTTP(w, r)
}
