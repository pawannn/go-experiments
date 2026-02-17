package main

import "fmt"

func SliceToChan(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()

	return out
}

func sqNums(ch <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range ch {
			out <- num * num
		}
		close(out)
	}()
	return out
}

func StartPipeline() {
	nums := []int{1, 2, 3, 4, 5, 6}
	sliceChan := SliceToChan(nums)
	sqChan := sqNums(sliceChan)

	for num := range sqChan {
		fmt.Println(num)
	}
}
