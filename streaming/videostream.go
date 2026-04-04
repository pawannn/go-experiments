package main

import (
	"net/http"
	"os"
)

func startVideoStreaming() {
	http.HandleFunc("/video", StreamVideo)
	http.ListenAndServe(":8080", nil)
}

func StreamVideo(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("video.mp4")
	if err != nil {
		http.Error(w, "File not found", 404)
		return
	}
	defer file.Close()

	stat, _ := file.Stat()

	http.ServeContent(w, r, stat.Name(), stat.ModTime(), file)
}
