package main

import (
	"fmt"
	"sync"
)

// https://leetcode.com/problems/building-h2o/description/?envType=company&envId=amazon&favoriteSlug=amazon-all

type H2O struct {
	mu     sync.Mutex
	cond   *sync.Cond
	hCount int
	wg     sync.WaitGroup
}

func newH2O() *H2O {
	hho := &H2O{}
	hho.cond = sync.NewCond(&hho.mu)
	return hho
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

	if h.hCount < 2 {
		h.cond.Wait()
	}

	fmt.Println("O")
	h.cond.Broadcast()
}

func GenerateWater() {
	h2o := newH2O()

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
