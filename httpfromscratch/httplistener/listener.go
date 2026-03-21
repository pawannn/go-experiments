package main

import (
	"fmt"
	"httpfromscratch/internal/request"
	"httpfromscratch/internal/response"
	"httpfromscratch/internal/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const port uint16 = 8080

func main() {
	fmt.Println("Server listening at port : ", port)
	s, err := server.Serve(port, func(w *response.ResponseWriter, r *request.Request) {
		statusCode := response.StatusOk
		body := []byte{}

		switch r.RequestLine.RequestTarget {
		case "/test":
			statusCode = response.StatusBadRequest
			body = []byte("bad request")
		case "/test1":
			statusCode = response.StatusInternalServerError
			body = []byte("internal server error")
		case "/test2":
			statusCode = response.StatusOk
			body = []byte("request accepted")
		}

		w.SendResponse(statusCode, body)
	})

	if err != nil {
		log.Fatal(err)
	}

	defer s.Close()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
	<-sigChan
	fmt.Println("Shutting down server gracefully...")
}
