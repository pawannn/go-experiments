package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/stream", StreamEvents)
	http.ListenAndServe(":8080", nil)
}

func StreamEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	str := "pawan kalyan is good boy"

	for part := range strings.SplitSeq(str, " ") {
		fmt.Fprintf(w, "data: %s\n\n", part)
		flusher.Flush()
		time.Sleep(time.Millisecond * 800)
	}
}
