package main

import (
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	rl := NewRateLimiter(5, time.Second*5)

	http.Handle("/test", rl.Middleware(http.HandlerFunc(Handler)))
	http.ListenAndServe(":8080", nil)
}
