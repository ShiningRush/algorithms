package main

import "fmt"

func longestPalindrome(s string) string {
	ps := prehandle(s)
	rl := make([]int, len(ps), len(ps))

	for i := 0; i < len(ps); i++ {
		ss := findPalindromicString(ps, i)
		rl[i] = len(ss) / 2
	}

	maxIdx, maxVal := 0, rl[0]
	for i, v := range rl {
		if v > maxVal {
			maxIdx = i
			maxVal = rl[i]
		}
	}

	si := (maxIdx - maxVal) / 2
	return s[si : si+maxVal]
}

func findPalindromicString(s string, idx int) string {
	if idx == 0 {
		return ""
	}

	if idx == len(s)-1 {
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

func prehandle(s string) string {
	ns := []rune{'#'}
	for _, v := range s {
		ns = append(ns, v)
		ns = append(ns, '#')
	}
	return string(ns)
}

func main() {
	fmt.Println(prehandle("ccd"))
	fmt.Println(longestPalindrome("ccd"))
}
