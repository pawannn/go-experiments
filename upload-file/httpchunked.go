package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func StartHttpChunkedUpload() {
	http.HandleFunc("/file", handleHttpFileUpload)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleHttpFileUpload(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20) // ~ 10 MB
	if err := r.ParseMultipartForm(5 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.MultipartForm.RemoveAll()

	uf, header, err := r.FormFile("media")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fileName := header.Filename

	f, err := os.Create(fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := io.Copy(f, uf); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("uploaded successfully"))
}
