package main

import (
	"fmt"
	"sync"
)

type OddEven struct {
	mu   sync.Mutex
	cond *sync.Cond
	i    int
}

func (oE *OddEven) Odd(wg *sync.WaitGroup, limit int) {
	defer wg.Done()

	for {
		oE.mu.Lock()

		if oE.i >= limit {
			oE.cond.Broadcast()
			oE.mu.Unlock()
			return
		}

		for oE.i%2 == 0 {
			oE.cond.Wait()
		}

		fmt.Println("Odd : ", oE.i)
		oE.i++

		oE.cond.Broadcast()
		oE.mu.Unlock()
	}
}

func (oE *OddEven) Even(wg *sync.WaitGroup, limit int) {
	defer wg.Done()

	for {
		oE.mu.Lock()

		if oE.i >= limit {
			oE.cond.Broadcast()
			oE.mu.Unlock()
			return
		}

		for oE.i%2 != 0 {
			oE.cond.Wait()
		}

		fmt.Println("Even : ", oE.i)
		oE.i++

		oE.cond.Broadcast()
		oE.mu.Unlock()
	}
}

func startOddEven() {
	oE := OddEven{}
	oE.cond = sync.NewCond(&oE.mu)

	limit := 9

	var wg sync.WaitGroup
	wg.Add(2)

	go oE.Even(&wg, limit)
	go oE.Odd(&wg, limit)

	wg.Wait()
}
