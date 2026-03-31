package main

// https://leetcode.com/problems/roman-to-integer
func romanToInt(s string) int {
	if s == "" {
		return 0
	}

	res := 0
	romans := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	res += romans[s[len(s)-1]]

	for i := len(s) - 2; i >= 0; i-- {
		if romans[s[i]] < romans[s[i+1]] {
			res -= romans[s[i]]
		} else {
			res += romans[s[i]]
		}
	}

	return res
}
