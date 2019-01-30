package main

import "fmt"

func main() {
	aArray := []int{1, 5, 7, 8, 12, 6, 3, 19, 26, 17}

	fmt.Printf("%+v", insertSort(aArray, 0, len(aArray)-1))
}

func insertSort(sortArray []int, start, end int) []int {
	for i := start + 1; i <= end; i++ {
		insertToArray(sortArray, sortArray[i], start, i)
	}

	return sortArray
}

func insertToArray(sortArray []int, insertValue, start, end int) {
	for i := start; i < end; i++ {
		if insertValue < sortArray[i] {
			rebuildRestArray(sortArray, i, end)
			sortArray[i] = insertValue
			break
		}
	}
}

func rebuildRestArray(bucket []int, startIdx, end int) {
	for idx := end; idx > startIdx; idx-- {
		temp := bucket[idx]
		bucket[idx] = bucket[idx-1]
		bucket[idx-1] = temp
	}
}
