package main

import (
	"net/http"
	"sync"
	"time"
)

type ResponseRecorder struct {
	http.ResponseWriter
	statusCode int
}

type state int

const (
	CLOSED = iota
	OPEN
	HALF_OPEN
)

type CircuitBreaker struct {
	mu sync.Mutex

	state state

	successThreshold int
	failureThreshold int

	successCount int
	failureCount int

	timeout time.Duration
	nextTry time.Time
}

func New(successThreshold int, failureThreshold int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		successThreshold: successThreshold,
		failureThreshold: failureThreshold,
		timeout:          timeout,
		state:            CLOSED,
	}
}

func (c *CircuitBreaker) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.mu.Lock()
		if c.state == OPEN {
			if c.nextTry.Before(time.Now()) {
				c.state = HALF_OPEN
			} else {
				http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
				c.mu.Unlock()
				return
			}
		}

		c.mu.Unlock()

		respRec := &ResponseRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(respRec, r)

		c.mu.Lock()
		defer c.mu.Unlock()

		if respRec.statusCode >= 500 {
			c.OnFailure()
		} else {
			c.OnSuccess()
		}
	})
}

func (c *CircuitBreaker) OnFailure() {
	c.failureCount++
	c.successCount = 0

	if c.failureCount >= c.failureThreshold {
		c.state = OPEN
		c.nextTry = time.Now().Add(c.timeout)
	}
}

func (c *CircuitBreaker) OnSuccess() {
	if c.state == HALF_OPEN {
		c.successCount++
		if c.successCount >= c.successThreshold {
			c.Reset()
		}
	} else {
		c.Reset()
	}
}

func (c *CircuitBreaker) Reset() {
	c.successCount = 0
	c.failureCount = 0

	c.state = CLOSED
}

func (r *ResponseRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}
