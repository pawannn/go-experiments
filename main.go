package main

import (
	"fmt"
	"sync"
)

type H2O struct {
	mu     sync.Mutex
	cond   *sync.Cond
	hCount int
	wg     sync.WaitGroup
}

func NewH2O() *H2O {
	h := &H2O{}
	h.cond = sync.NewCond(&h.mu)
	return h
}

func (h *H2O) Hydrogen() {
	defer h.wg.Done()

	h.mu.Lock()
	defer h.mu.Unlock()

	for h.hCount == 2 {
		h.cond.Wait()
	}

	fmt.Print("H")
	h.hCount++

	if h.hCount == 2 {
		h.cond.Broadcast()
	}
}

func (h *H2O) Oxygen() {
	defer h.wg.Done()

	h.mu.Lock()
	defer h.mu.Unlock()

	for h.hCount < 2 {
		h.cond.Wait()
	}

	fmt.Print("O")

	h.hCount = 0
	h.cond.Broadcast()
}

func main() {
	h2o := NewH2O()

	input := []rune("HOHOHH")

	for _, c := range input {
		h2o.wg.Add(1)
		if c == 'H' {
			go h2o.Hydrogen()
		} else {
			go h2o.Oxygen()
		}
	}

	h2o.wg.Wait()
	fmt.Println()
}
