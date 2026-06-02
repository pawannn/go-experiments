package main

func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]bool)
	maxLen := 0
	left := 0

	for right := 0; right < len(s); right++ {
		for seen[s[right]] {
			delete(seen, s[right])
			left++
		}

		seen[s[right]] = true

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}
