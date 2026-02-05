package main

import (
	"net/url"
	"sync"
)

type Backend struct {
	url   *url.URL
	alive bool
	mu    sync.Mutex
}

func NewBackend(addr string) (*Backend, error) {
	url, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}

	return &Backend{
		url:   url,
		alive: true,
	}, nil
}

func (b *Backend) SetAlive(alive bool) {
	b.mu.Lock()
	b.alive = alive
	b.mu.Unlock()
}

func (b *Backend) IsAlive() bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.alive
}
