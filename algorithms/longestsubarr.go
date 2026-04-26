package main

import "fmt"

func LongestSubarr(input []int, k int) int {
	longestSubarr := 0

	for i := 0; i < len(input)-1; i++ {
		sum := input[i]
		if input[i] == k {
			longestSubarr += 1
		}

		for j := i + 1; j < len(input); j++ {
			sum += input[j]
			if sum == k {
				longestSubarr += j - i
				break
			}
		}
	}

	return longestSubarr
}

func CheckLongestSubstr() {
	arr := []int{10, -10, 20, 30}
	k := 5

	longerSubArr := LongestSubarr(arr, k)

	fmt.Println(longerSubArr)
}
