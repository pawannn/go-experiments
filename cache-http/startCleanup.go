package main

import (
	"log"
	"time"
)

func (c *Cache) StartCleanUp() {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("Started Cleanup")
		now := time.Now().UnixNano()

		c.mu.Lock()

		for k, v := range c.Item {
			if v.Expiry < now {
				delete(c.Item, k)
			}
		}

		c.mu.Unlock()
	}
}
