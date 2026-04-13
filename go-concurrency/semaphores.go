package main

import (
	"fmt"
	"time"
)

// A semaphore in Go is a way to control how many goroutines run concurrently, usually implemented using channels.

func tradationalWorker(i int, sem chan struct{}) {
	sem <- struct{}{}

	fmt.Println("worker ", i, " started working")
	// Do some work

	time.Sleep(time.Second * 2)

	<-sem
}

func tradationalSemaphores() {
	sem := make(chan struct{}, 3) // Allow only 2 go routines to work at a time

	for i := range 5 {
		go tradationalWorker(i, sem)
	}

	time.Sleep(10 * time.Second)

}
