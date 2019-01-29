package main

import "fmt"

func main() {
	aArray := []int{1, 5, 7, 8, 12, 6, 3, 19, 26, 17}

	fmt.Printf("%+v", countingSort(aArray, 0, len(aArray)-1))
}

func countingSort(sortArray []int, start, end int) []int {
	min, max := findMinAndMax(sortArray, start, end)

	extraArray := make([]int, max-min+1, max-min+1)
	populateExtraArray(extraArray, sortArray, min)
	return buildResult(extraArray, min)
}

func findMinAndMax(sortArray []int, start, end int) (int, int) {
	min, max := sortArray[start], sortArray[start]

	for i := start + 1; i <= end; i++ {
		if sortArray[i] < min {
			min = sortArray[i]
		}

		if sortArray[i] > max {
			max = sortArray[i]
		}
	}

	return min, max
}

func populateExtraArray(extraArray, sortArray []int, min int) {
	for _, v := range sortArray {
		extraArray[v-min]++
	}
}

func buildResult(extraArray []int, min int) []int {
	var rs []int
	for k, v := range extraArray {
		for i := 0; i < v; i++ {
			rs = append(rs, min+k)
		}

	}

	return rs
}
