package main

import (
	"fmt"
)

func printOdd(ch chan int, done chan bool) {
	for i := 1; i <= 10; i += 2 {
		<-ch
		fmt.Println("Odd:", i)
		ch <- 1
	}
	done <- true
}

func printEven(ch chan int, done chan bool) {
	for i := 2; i <= 10; i += 2 {
		<-ch
		fmt.Println("Even:", i)
		ch <- 1
	}
	done <- true
}

func startOddEvenChan() {
	ch := make(chan int)
	done := make(chan bool)

	go printOdd(ch, done)
	go printEven(ch, done)

	ch <- 1

	<-done
	<-done
}
