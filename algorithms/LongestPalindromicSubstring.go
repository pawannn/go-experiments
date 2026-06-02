package main

func longestPalindrome(s string) string {
	res := ""
	maxLen := 0

	for i := 0; i < len(s); i++ {

		r, l := i, i
		for l >= 0 && r < len(s) && s[r] == s[l] {
			if r-l+1 > maxLen {
				maxLen = r - l + 1
				res = string(s[l : r+1])
			}
			l--
			r++
		}

		r, l = i, i+1
		for l >= 0 && r < len(s) && s[r] == s[l] {
			if r-l+1 > maxLen {
				maxLen = r - l + 1
				res = string(s[l : r+1])
			}
			l--
			r++
		}
	}

	return res
}
