package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type NewCacheRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	TTL   int64  `json:"ttl"`
}

func (s *Server) HandleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.HandleGet(w, r)
	case http.MethodPut:
		s.HandleSet(w, r)
	case http.MethodDelete:
		// Do Something
	}
}

func (s *Server) HandleSet(w http.ResponseWriter, r *http.Request) {
	var payload NewCacheRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unable to parse payload"))
		return
	}

	s.c.Set(payload.Key, payload.Value, time.Duration(payload.TTL)*time.Second)

	w.WriteHeader(http.StatusOK)
}

func (s *Server) HandleGet(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	key := query.Get("key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("key query is required"))
		return
	}

	value, exist := s.c.Get(key)
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("cache with the key not found"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(value))
}

func (s *Server) HandleDelete(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	key := query.Get("key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("key query is required"))
		return
	}

	s.c.Delete(key)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("cache deleted successfully"))
}
