package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const uploadDir = "./upload/"

func HandleChunk(w http.ResponseWriter, r *http.Request) {
	fileID := r.FormValue("fileID")
	if strings.TrimSpace(fileID) == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("provide a valid fileID"))
		return
	}

	chunkID, err := strconv.Atoi(r.FormValue("chunkID"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("provide a valid chunk ID"))
		return
	}

	chunkSize, err := strconv.Atoi(r.FormValue("chunkSize"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("provide a valid chunk size"))
		return
	}

	totalLength, err := strconv.Atoi(r.FormValue("totalLength"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("provide a valid total length"))
		return
	}

	chunkData, _, err := r.FormFile("chunk")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid chunk data"))
		return
	}

	dest := uploadDir + fileID
	dst, err := os.OpenFile(dest, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to save chunk data"))
		return
	}
	if chunkID == 0 {
		dst.Truncate(int64(totalLength))
	}
	defer dst.Close()

	offset := int64(chunkID) * int64(chunkSize)

	_, err = dst.Seek(offset, 0)
	if err != nil {
		http.Error(w, "seek failed", http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(dst, chunkData)
	if err != nil {
		http.Error(w, "copy failed", http.StatusInternalServerError)
		return
	}

	fmt.Printf("FileID: Received chunk %d\n", chunkID)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("chunk uploaded successfully"))
}

func HandleComplete(w http.ResponseWriter, r *http.Request) {
	fileID := r.FormValue("fileID")
	if strings.TrimSpace(fileID) == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("provide a valid fileID"))
		return
	}

	filePath := uploadDir + fileID

	info, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "file not found", http.StatusNotFound)
			return
		}
		http.Error(w, "unable to access file", http.StatusInternalServerError)
		return
	}

	actualSize := info.Size()

	expectedSizeStr := r.FormValue("totalLength")
	expectedSize, err := strconv.ParseInt(expectedSizeStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid totalLength", http.StatusBadRequest)
		return
	}

	if actualSize != expectedSize {
		http.Error(w, fmt.Sprintf("file size mismatch: expected %d, got %d", expectedSize, actualSize), http.StatusBadRequest)
		return
	}

	fmt.Printf("FileID: Upload complete (%s), size verified: %d bytes\n", fileID, actualSize)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("upload completed successfully"))
}
