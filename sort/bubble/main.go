package main

import (
	"fmt"
)

func main() {
	aArray := []int{1, 5, 7, 8, 12, 6, 3, 19, 26, 17}
	bubbleSort(aArray)
	fmt.Printf("%+v", aArray)
}

func bubbleSort(sortArray []int) {
	bubble(sortArray, len(sortArray))
}

func bubble(sortArray []int, endIdx int) {
	hasSwap := false

	for i := 0; i < endIdx; i++ {
		if i+1 < endIdx && sortArray[i] > sortArray[i+1] {
			swap(sortArray, i, i+1)
			hasSwap = true
		}
	}

	if hasSwap && endIdx >= 2 {
		bubble(sortArray, endIdx-1)
	}
}

func swap(sortArray []int, i int, j int) {
	temp := sortArray[i]
	sortArray[i] = sortArray[j]
	sortArray[j] = temp
}
