package main

import (
	"fmt"
	"runtime"
	"sync"
)

type Item struct {
	sum  int
	pair [2]int
}

func GetUnqiueSumWithElements(nums []int) map[int][][2]int {
	var wg sync.WaitGroup
	pairsChan := make(chan [2]int)
	sumChan := make(chan Item)

	concurrency := runtime.NumCPU()
	for range concurrency {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for p := range pairsChan {
				sumChan <- Item{
					sum:  p[0] + p[1],
					pair: p,
				}
			}
		}()
	}

	go func() {
		for i := 0; i < len(nums); i++ {
			for j := i + 1; j < len(nums); j++ {
				pairsChan <- [2]int{nums[i], nums[j]}
			}
		}
		close(pairsChan)
	}()

	go func() {
		wg.Wait()
		close(sumChan)
	}()

	result := make(map[int][][2]int)
	for s := range sumChan {
		result[s.sum] = append(result[s.sum], s.pair)
	}

	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sumWithElements := GetUnqiueSumWithElements(nums)

	for key, val := range sumWithElements {
		fmt.Printf("%d : %v\n", key, val)
	}
}
