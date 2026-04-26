package main

import "fmt"

// https://leetcode.com/problems/valid-parentheses
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := []rune{}

	for _, ch := range s {
		fmt.Println(ch)
		switch ch {
		case '(', '{', '[':
			stack = append(stack, ch)

		case ')', '}', ']':
			if len(stack) == 0 {
				return false
			}

			lastBracket := stack[len(stack)-1]

			if ch == ')' && lastBracket == '(' ||
				ch == ']' && lastBracket == '[' ||
				ch == '}' && lastBracket == '{' {
				stack = stack[0 : len(stack)-1]
			}
		}
	}

	return len(stack) == 0
}
