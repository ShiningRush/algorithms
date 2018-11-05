package main

import "fmt"

func main() {
	aArray := []int{1, 5, 7, 8, 12, 6, 3, 19, 26, 17}

	fmt.Printf("%+v", mergeSort(aArray, 0, len(aArray)-1))
}

func mergeSort(sortArray []int, start, end int) []int {
	if start == end {
		return sortArray[start : start+1]
	}

	if end == start+1 {
		return merge(sortArray[start:start+1], sortArray[end:end+1])
	}

	mid := (start + end) / 2
	rs := merge(mergeSort(sortArray, start, mid), mergeSort(sortArray, mid+1, end))
	return rs
}

func merge(a, b []int) []int {
	var rs []int
	var aFlg, bFlg bool
	count := len(a) + len(b)
	for i := 0; i < count; i++ {
		if (a[0] < b[0] && !aFlg) || bFlg {
			rs = append(rs, a[0])
			if len(a) == 1 {
				aFlg = true
			} else {
				a = a[1:]
			}
		} else {
			rs = append(rs, b[0])
			if len(b) == 1 {
				bFlg = true
			} else {
				b = b[1:]
			}
		}
	}

	return rs
}
