package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

type SetPayload struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Value struct {
	Value  string
	Expiry time.Time
}

type Cache struct {
	mu   sync.Mutex
	data map[string]Value
}

func (c *Cache) Set(key string, value string, expiry time.Time) error {
	if expiry.Before(time.Now()) {
		return fmt.Errorf("EXPIRY should be in future")
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	cacheVal := Value{
		Value:  value,
		Expiry: expiry,
	}

	c.data[key] = cacheVal

	return nil
}

func (c *Cache) Get(key string) string {
	c.mu.Lock()
	defer c.mu.Unlock()

	val, ok := c.data[key]
	if !ok {
		return ""
	}

	if val.Expiry.Before(time.Now()) {
		delete(c.data, key)
		return ""
	}

	return val.Value
}

func (c *Cache) Has(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, ok := c.data[key]

	if !ok {
		return false
	} else {
		return true
	}
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, ok := c.data[key]
	if ok {
		delete(c.data, key)
	}
}

func (c *Cache) SetKey(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid request body"))
		return
	}

	var payload SetPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid payload"))
		return
	}

	if strings.TrimSpace(payload.Key) == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid cache key"))
		return
	}

	if err := c.Set(payload.Key, payload.Value, time.Now().Add(time.Second*30)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to create a data  at the moment, try again later"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("data added to cache successfully"))
}

func (c *Cache) GetKey(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	query := r.URL.Query()
	key := query.Get("key")

	if strings.TrimSpace(key) == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid cache key"))
		return
	}

	val := c.Get(key)

	response := map[string]string{
		"key": key,
		"val": val,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func startCacheServer() {
	c := Cache{
		data: make(map[string]Value),
	}

	http.HandleFunc("/set", c.SetKey)
	http.HandleFunc("/get", c.GetKey)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("server failed:", err)
	}
}
