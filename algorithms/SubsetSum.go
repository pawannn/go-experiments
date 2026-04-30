package main

import (
	"fmt"
	"sort"
	"sync"
)

// Get unique sums of an array, using go routines and print in a synchronized way
type SubsetSum struct {
	arr []int
	sum map[int]struct{}
	mu  sync.Mutex
	wg  sync.WaitGroup
}

func (s *SubsetSum) Helper(index, currentSum int) {
	defer s.wg.Done()

	if index == len(s.arr) {
		s.mu.Lock()
		s.sum[currentSum] = struct{}{}
		s.mu.Unlock()
		return
	}

	s.wg.Add(1)
	go s.Helper(index+1, currentSum+s.arr[index])

	s.wg.Add(1)
	go s.Helper(index+1, currentSum)
}

func (s *SubsetSum) Solve() {
	s.wg.Add(1)
	go s.Helper(0, 0)
	s.wg.Wait()
}

func CheckSubSetSum() {
	s := SubsetSum{
		arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		sum: make(map[int]struct{}),
	}

	s.Solve()

	result := []int{}
	for i := range s.sum {
		result = append(result, i)
	}

	sort.Ints(result)

	fmt.Println(result)
}
