package main

import "fmt"

func main() {
	aArray := []int{1, 5, 7, 8, 12, 6, 3, 19, 26, 17}

	fmt.Printf("%+v", bucketSort(aArray, 0, len(aArray)-1))
}

func bucketSort(sortArray []int, start, end int) []int {
	var buckets [10][]int

	populateBuckets(&buckets, sortArray)
	return buildResult(buckets)
}

func populateBuckets(buckets *[10][]int, sortArray []int) {
	for _, v := range sortArray {
		buckets[v/5] = insertToBucket(buckets[v/5], v)
	}
}

func insertToBucket(bucket []int, insertValue int) []int {
	bucket = append(bucket, insertValue)
	for k, v := range bucket {
		if v > insertValue {
			rebuildRestArray(bucket, k)
			bucket[k] = insertValue
			break
		}
	}

	return bucket
}

func rebuildRestArray(bucket []int, startIdx int) {
	for idx := len(bucket) - 1; idx > startIdx; idx-- {
		temp := bucket[idx]
		bucket[idx] = bucket[idx-1]
		bucket[idx-1] = temp
	}
}

func buildResult(buckets [10][]int) []int {
	var rs []int
	for _, bucket := range buckets {
		for _, v := range bucket {
			rs = append(rs, v)
		}
	}

	return rs
}
