package main

import "fmt"

func longestPalindrome(s string) string {
	rs := ""
	maxLength := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			subs := s[i:j]
			if len(subs) > maxLength && checkIsPalindromicString(subs) && len(subs) > maxLength {
				rs = subs
				maxLength = len(subs)
			}
		}
	}
	return rs
}

func checkIsPalindromicString(s string) bool {
	if len(s) == 1 {
		return true
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(checkIsPalindromicString("ccd"))
	fmt.Println(checkIsPalindromicString("cc"))
	fmt.Println(checkIsPalindromicString("cbc"))
	fmt.Println(longestPalindrome("ccd"))
}
