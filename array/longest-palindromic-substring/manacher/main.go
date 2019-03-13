package main

import "fmt"

func longestPalindrome(s string) string {
	ps := prehandle(s)
	rl := make([]int, len(s), len(s))

	return ""
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
