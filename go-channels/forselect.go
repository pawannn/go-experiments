package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func ForSelect(charChan <-chan string, shutdownChan <-chan os.Signal) {
	for {
		select {
		case msg := <-charChan:
			fmt.Println("Got message : ", msg)
		case <-shutdownChan:
			fmt.Println("Bye...")
			return
		default:
			fmt.Println("Nothing to do")
		}
	}
}

func StartForSelectRoutine() {
	charChan := make(chan string)
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGTERM, syscall.SIGINT)

	go ForSelect(charChan, shutdown)

	for {
		charChan <- "hello"
	}
}
