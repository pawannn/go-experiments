package main

import (
	"fmt"
	"runtime"
	"sort"
	"sync"
)

func GetUnqueSums(nums []int) []int {
	var wg sync.WaitGroup

	pairChan := make(chan [2]int)
	sumChan := make(chan int)

	concurrency := runtime.NumCPU()

	for range concurrency {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for p := range pairChan {
				sumChan <- p[0] + p[1]
			}
		}()
	}

	go func() {
		for i := 0; i < len(nums); i++ {
			for j := i + 1; j < len(nums); j++ {
				pairChan <- [2]int{nums[i], nums[j]}
			}
		}
		defer close(pairChan)
	}()

	go func() {
		wg.Wait()
		close(sumChan)
	}()

	unique := make(map[int]struct{})
	for s := range sumChan {
		unique[s] = struct{}{}
	}

	resultSet := make([]int, 0, len(unique))
	for key := range unique {
		resultSet = append(resultSet, key)
	}

	sort.Ints(resultSet)

	return resultSet
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sums := GetUnqueSums(nums)
	fmt.Println(sums)
}
