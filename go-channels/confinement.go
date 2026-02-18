package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

func processData(res *int, data int) {
	num := data * 2
	lock.Lock()
	defer lock.Unlock()
	*res = num
}

func StartConfinement() {
	var wg sync.WaitGroup

	input := []int{1, 2, 3, 4, 5, 6}
	result := make([]int, len(input))

	for i, data := range input {
		wg.Go(func() {
			processData(&result[i], data)
		})
	}

	wg.Wait()

	fmt.Println(result)
}
