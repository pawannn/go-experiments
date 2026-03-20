package main

import (
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/pawannn/httpfromscratch/internal/request"
	"github.com/pawannn/httpfromscratch/internal/response"
	"github.com/pawannn/httpfromscratch/internal/server"
)

const port uint16 = 8080

func main() {
	s, err := server.Serve(port, func(w io.Writer, req *request.Request) *response.HandlerError {
		switch req.RequestLine.RequestTarget {
		case "/yourporblem":
			return &response.HandlerError{
				Code:    response.StatusBadRequest,
				Message: "Your problem not my problem",
			}
		case "/myproblem":
			return &response.HandlerError{
				Code:    response.StatusInternalServerError,
				Message: "my problem not your problem",
			}
		case "/use-neovim":
			w.Write([]byte("all good\n"))
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Error starting server : %v", err)
	}

	defer s.Close()
	log.Println("Server started listening at port : ", port)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Shutting down gracefully")
}
