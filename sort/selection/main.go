package main

import "fmt"

func main() {
	aArray := []int{1, 5, 7, 8, 12, 6, 3, 19, 26, 17}
	selectSort(aArray, 0, len(aArray))
	fmt.Printf("%+v", aArray)
}

func selectSort(sortArray []int, start, end int) {
	if start >= end {
		return
	}

	max := sortArray[start]
	maxIdx := start
	for i := start + 1; i < end; i++ {
		if sortArray[i] > max {
			maxIdx = i
			max = sortArray[i]
		}
	}

	swap(sortArray, maxIdx, end-1)
	selectSort(sortArray, start, end-1)
}

func swap(sortArray []int, i, j int) {
	temp := sortArray[i]
	sortArray[i] = sortArray[j]
	sortArray[j] = temp
}
