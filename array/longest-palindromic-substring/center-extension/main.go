package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}

	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		} else {
			return s[0:1]
		}
	}
	var longest string
	maxCount := 0
	for i := 0; i < len(s)-1; i++ {
		s1 := findPalindromicString(s, i)
		if len(s1) > maxCount {
			longest = s1
			maxCount = len(s1)
		}

		s2 := findPalindromicString2(s, i)
		if len(s2) > maxCount {
			longest = s2
			maxCount = len(s2)
		}
	}

	return longest
}

func findPalindromicString(s string, idx int) string {
	if idx == 0 {
		return ""
	}

	for i, j := idx-1, idx+1; ; {
		if s[i] != s[j] {
			return s[i+1 : j]
		}

		i, j = i-1, j+1
		if i < 0 || j >= len(s) {
			return s[i+1 : j]
		}
	}
}

func findPalindromicString2(s string, idx int) string {

	for i, j := idx, idx+1; ; {
		if s[i] != s[j] {
			return s[i+1 : j]
		}

		i, j = i-1, j+1
		if i < 0 || j >= len(s) {
			return s[i+1 : j]
		}
	}
}

func main() {
	fmt.Println(longestPalindrome("ccd"))
}
