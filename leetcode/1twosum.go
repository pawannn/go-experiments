package main

// https://leetcode.com/problems/two-sum
func twoSum(nums []int, target int) []int {
	read := make(map[int]int)

	for idx, num := range nums {
		diff := target - num
		i, ok := read[diff]
		if !ok {
			read[num] = idx
			continue
		}

		return []int{i, idx}
	}

	return []int{}
}
