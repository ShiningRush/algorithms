package main

func main() {
	nums1 := []int{3, 4, 6, 5}
	nums2 := []int{9, 1, 2, 5, 8, 3}
	maxNumber(nums1, nums2, 5)
}

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	var rs []int
	var idx []int
	for i := 0; i < k; i++ {
		rs = append(rs, findMaxNumer(nums1, nums2, &idx))
	}

	return rs[0:]
}

func findMaxNumer(s1, s2 []int, sIdx *[]int) int {
	max := s1[0]
	maxIdx := 0
	for i := 0; i < len(s1)+len(s2); i++ {
		if i < len(s1) {
			if s1[i] > max && !isExtracted(i, *sIdx) {
				max = s1[i]
				maxIdx = i
			}
		} else {
			if s2[i-len(s1)] > max && !isExtracted(i, *sIdx) {
				max = s2[i-len(s1)]
				maxIdx = i
			}
		}
	}

	*sIdx = append(*sIdx, maxIdx)

	return max
}

func isExtracted(idx int, sIdx []int) bool {
	for _, idxV := range sIdx {
		if idx == idxV {
			return true
		}
	}

	return false
}
