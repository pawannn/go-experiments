package main

import (
	"net/url"
	"sync"
)

type Backend struct {
	URL         *url.URL
	Alive       bool
	ActiveConns int64
	mu          sync.RWMutex
}

func NewBackend(addr string) (*Backend, error) {
	url, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}

	return &Backend{
		URL:   url,
		Alive: true,
	}, nil
}

func (b *Backend) SetAlive(alive bool) {
	b.mu.Lock()
	b.Alive = alive
	b.mu.Unlock()
}

func (b *Backend) IsAlive() bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.Alive
}
