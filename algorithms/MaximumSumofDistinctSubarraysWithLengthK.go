package main

func maximumSubarraySum(nums []int, k int) int64 {
	seen := make(map[int]bool)
	left := 0
	var windowSum int64 = 0
	var maxSum int64 = 0

	for right := 0; right < len(nums); right++ {
		for seen[nums[right]] {
			delete(seen, nums[left])
			windowSum -= int64(nums[left])
			left++
		}

		seen[nums[right]] = true
		windowSum += int64(nums[right])

		if right-left+1 == k {
			if windowSum > maxSum {
				maxSum = windowSum
			}

			delete(seen, nums[left])
			windowSum -= int64(nums[left])
			left++
		}
	}

	return int64(maxSum)
}
