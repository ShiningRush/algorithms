package main

import "fmt"

func main() {
	aArray := []int{123, 52, 7, 8, 12, 6, 3, 19, 26, 17, 4, 99, 5}

	fmt.Printf("%+v", heapSort(aArray, 0, len(aArray)-1))
}

func heapSort(sortArray []int, start, end int) []int {
	for i := len(sortArray) - 1; i >= 0; i-- {
		heapify(sortArray, i)
		swap(sortArray, 0, i)
	}
	return sortArray
}

// based siftdown
func heapify(array []int, end int) {
	for i := 0; i <= end; i++ {
		siftDown(array, i)
	}
}

func siftDown(heap []int, idx int) {
	if idx == 0 {
		return
	}

	parentIdx := 0
	// left child
	if idx%2 == 1 {
		parentIdx = (idx - 1) / 2
	}

	// right child
	if idx%2 == 0 {
		parentIdx = (idx - 2) / 2
	}

	if heap[parentIdx] < heap[idx] {
		swap(heap, parentIdx, idx)
	}

	siftDown(heap, parentIdx)
}

func swap(array []int, i, j int) {
	temp := array[i]
	array[i] = array[j]
	array[j] = temp
}
