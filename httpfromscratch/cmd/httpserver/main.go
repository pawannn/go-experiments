package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/pawannn/httpfromscratch/internal/server"
)

const port uint16 = 8080

func main() {
	s, err := server.Serve(port)
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
