package main

import (
	"net/http"
	"sync"
	"time"
)

type Client struct {
	tokens    int
	lastRefil time.Time
}

func NewClient(tokens int) *Client {
	return &Client{
		tokens:    tokens,
		lastRefil: time.Now(),
	}
}

type RateLimiter struct {
	mu        sync.Mutex
	clients   map[string]*Client
	maxTokens int
	refilRate time.Duration
}

func NewRateLimiter(maxTokens int, refilRate time.Duration) *RateLimiter {
	rL := &RateLimiter{
		clients:   make(map[string]*Client),
		maxTokens: maxTokens,
		refilRate: refilRate,
	}

	go rL.Cleanup()

	return rL
}

func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr

		if !rl.Allow(ip) {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Too many requests, please try again later"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (rl *RateLimiter) Allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	c, exits := rl.clients[ip]
	if !exits {
		newClient := NewClient(rl.maxTokens - 1)
		rl.clients[ip] = newClient
		return true
	}

	timeElapsed := time.Since(c.lastRefil)

	refilTokens := int(timeElapsed / rl.refilRate)

	if refilTokens > 0 {
		c.tokens = refilTokens

		c.tokens = min(c.tokens, rl.maxTokens)

		c.lastRefil = time.Now()
	}

	if c.tokens <= 0 {
		return false
	}

	c.tokens--

	return true
}

func (rL *RateLimiter) Cleanup() {
	ticker := time.NewTicker(time.Minute * 5)

	for range ticker.C {
		rL.mu.Lock()
		for ip, c := range rL.clients {
			if time.Since(c.lastRefil) > 10*time.Minute {
				delete(rL.clients, ip)
			}
		}

		rL.mu.Unlock()
	}
}
