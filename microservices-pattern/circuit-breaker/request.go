package main

import (
	"math/rand/v2"
	"net/http"
	"time"
)

func UnstableHandler(w http.ResponseWriter, r *http.Request) {
	if rand.Float64() < 0.6 {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("done"))
}

func main() {
	cb := New(3, 2, time.Second*10)

	http.Handle("/pattern", cb.Middleware(http.HandlerFunc(UnstableHandler)))
	http.ListenAndServe(":8080", nil)
}
