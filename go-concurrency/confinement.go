// Confinement in Go concurrency means restricting access to shared data so that only one goroutine owns and modifies it, preventing race conditions without needing locks.

package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

func ManageTickets(done <-chan bool, ticketChan <-chan int, tickets *int) {
	for {
		select {
		case user := <-ticketChan:
			if *tickets > 0 {
				*tickets--
				fmt.Println("User : ", user, " tickets remaining : ", *tickets)
			} else {
				fmt.Println("User : ", user, " ticket not found")
			}
		case <-done:
			fmt.Println("tickets remaining : ", *tickets)
			return
		}
	}
}

func BuyTickets(wg *sync.WaitGroup, ticketChan chan int, userID int) {
	defer wg.Done()
	ticketChan <- userID
}

func main() {
	var wg sync.WaitGroup
	tickets := 500
	ticketsChan := make(chan int)
	doneChan := make(chan bool)

	go ManageTickets(doneChan, ticketsChan, &tickets)

	for userID := range 2000 {
		wg.Add(1)
		go BuyTickets(&wg, ticketsChan, userID)
	}

	wg.Wait()
}
