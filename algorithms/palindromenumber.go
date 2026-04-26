package main

import "strconv"

// https://leetcode.com/problems/palindrome-number
func isPalindrome(x int) bool {
	str := strconv.Itoa(x)

	i := 0
	j := len(str) - 1

	for i < j {
		if str[i] != str[j] {
			return false
		}

		i++
		j--
	}

	return true
}
