package main

import (
	"net/http"
	"net/http/httputil"
	"sync"
	"sync/atomic"
	"time"
)

type LoadBalancer struct {
	backends []*Backend
	current  uint64
	mu       sync.Mutex
}

func NewLoadBalancer(backends []*Backend) *LoadBalancer {
	return &LoadBalancer{
		backends: backends,
		current:  0,
	}
}

func (lb *LoadBalancer) NextBackend() *Backend {
	n := len(lb.backends)
	for i := 0; i < n; i++ {
		idx := atomic.AddUint64(&lb.current, 1) % uint64(n)
		b := lb.backends[idx]
		if b.IsAlive() {
			return b
		}
	}
	return nil
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	backend := lb.NextBackend()
	if backend == nil {
		http.Error(w, "No backend available", http.StatusServiceUnavailable)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(backend.URL)

	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		backend.SetAlive(false)
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
	}

	proxy.ServeHTTP(w, r)
}

func healthCheck(lb *LoadBalancer) {
	for {
		for _, b := range lb.backends {
			resp, err := http.Get(b.URL.String() + "/health")
			if err != nil || resp.StatusCode != 200 {
				b.SetAlive(false)
				continue
			}
			b.SetAlive(true)
		}
		time.Sleep(10 * time.Second)
	}
}
