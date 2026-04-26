package main

// https://leetcode.com/problems/longest-common-prefix
func longestCommonPrefix(strs []string) string {
	res := ""

	for i := 0; i < len(strs[0]); i++ {
		for _, s := range strs {
			if i == len(s) || s[i] != strs[0][i] {
				return res
			}
		}

		res += string(strs[0][i])
	}

	return res
}
