package main

import (
	"sync"
	"time"
)

type Item struct {
	Value  string
	Expiry int64
}

type Cache struct {
	Item map[string]Item
	mu   sync.RWMutex
}

func NewTTLCache() *Cache {
	c := &Cache{
		Item: make(map[string]Item),
	}

	go c.StartCleanUp()

	return c
}

func (c *Cache) Set(key string, value string, expiry time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	e := time.Now().Add(expiry).UnixNano()
	c.Item[key] = Item{
		Value:  value,
		Expiry: e,
	}
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, ok := c.Item[key]
	if !ok {
		return "", false
	}

	expiry := item.Expiry

	if time.Now().UnixNano() > expiry {
		return "", false
	}

	return item.Value, true
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.Item, key)
}
