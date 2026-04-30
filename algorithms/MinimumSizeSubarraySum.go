package main

// https://leetcode.com/problems/minimum-size-subarray-sum/
// Input: target = 7, nums = [2,3,1,2,4,3]
// Output: 2
func minSubArrayLen(target int, nums []int) int {
	sum := 0
	left := 0
	res := len(nums) + 1

	for right := range len(nums) {
		sum += nums[right]

		for sum >= target {
			if right-left+1 < res {
				res = right - left + 1
			}

			sum -= nums[left]
			left++
		}
	}

	if res == len(nums)+1 {
		return 0
	}

	return res
}
