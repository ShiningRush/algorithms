package main

import "fmt"

func main() {
	aArray := []int{1, 5, 7, 8, 12, 6, 3, 19, 26, 17}
	quickSort(aArray)
	fmt.Printf("%+v", aArray)
}

func quickSort(sortArray []int) {
	sort(sortArray, 0, len(sortArray)-1)
}

func sort(sortArray []int, startIdx int, endIdx int) {
	if startIdx >= endIdx {
		return
	}

	if endIdx == startIdx+1 {
		if sortArray[startIdx] > sortArray[endIdx] {
			swap(sortArray, startIdx, endIdx)
		}

		return
	}

	bv := sortArray[startIdx]
	j := startIdx + 1
	for i := startIdx + 1; i <= endIdx; i++ {
		if sortArray[i] < bv {
			swap(sortArray, i, j)
			j++
		}
	}

	swap(sortArray, startIdx, j-1)
	sort(sortArray, startIdx, j-2)
	sort(sortArray, j, endIdx)
}

func swap(sortArray []int, i int, j int) {
	temp := sortArray[i]
	sortArray[i] = sortArray[j]
	sortArray[j] = temp
}
