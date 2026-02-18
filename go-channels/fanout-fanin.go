package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

func RepeatFunc[T any, K any](done <-chan T, fn func() K) <-chan K {
	generateChan := make(chan K)
	go func() {
		defer close(generateChan)
		for {
			select {
			case <-done:
				return
			case generateChan <- fn():
			}
		}
	}()

	return generateChan
}

func PrimeFinder(done <-chan int, randNumChan <-chan int) <-chan int {
	isPrime := func(num int) bool {
		if num < 2 {
			return false
		}
		for i := 2; i*i < num; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}

	primeChannel := make(chan int)
	go func() {
		defer close(primeChannel)
		for {
			select {
			case <-done:
				return
			case num, ok := <-randNumChan:
				if !ok {
					return
				}
				if isPrime(num) {
					primeChannel <- num
				}
			}
		}
	}()

	return primeChannel
}

func Take[T any, K any](done <-chan T, in <-chan K, limit int) <-chan K {
	takeChan := make(chan K)
	go func() {
		defer close(takeChan)
		for range limit {
			select {
			case <-done:
				return
			case takeChan <- <-in:
			}
		}
	}()

	return takeChan
}

func FanIn[T any, K any](done <-chan T, inChans ...<-chan K) <-chan K {
	stream := make(chan K)
	var wg sync.WaitGroup

	readChannel := func(in <-chan K) {
		wg.Go(func() {
			for {
				select {
				case <-done:
					return
				case stream <- <-in:
				}
			}
		})
	}

	for _, ch := range inChans {
		go readChannel(ch)
	}

	go func() {
		wg.Wait()
		close(stream)
	}()

	return stream
}

func main() {
	done := make(chan int)
	randomGenerator := func() int { return rand.Intn(50000000) }

	randomNumChan := RepeatFunc(done, randomGenerator)

	cpuCount := runtime.NumCPU()

	primeFinderChannels := make([]<-chan int, cpuCount)
	for i := range cpuCount {
		primeFinderChannels[i] = PrimeFinder(done, randomNumChan)
	}

	fanInChannel := FanIn(done, primeFinderChannels...)

	for num := range Take(done, fanInChannel, 2) {
		fmt.Println(num)
	}
}
