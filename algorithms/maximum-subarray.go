package main

// https://leetcode.com/problems/maximum-subarray/
func maxSubArray(nums []int) int {
	sum := nums[0]
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		if sum+nums[i] > nums[i] {
			sum += nums[i]
		} else {
			sum = nums[i]
		}

		if sum > res {
			res = sum
		}
	}

	return res
}
