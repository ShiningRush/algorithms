package main

import (
	"fmt"
)

var rsArray [][]int

func superEggDrop(K int, N int) int {
	if rsArray == nil {
		initSlice(200, 10000)
	}

	if K == 1 {
		return N
	}

	if N <= 2 {
		return N
	}

	if N == 3 {
		return 2
	}

	if N == 4 {
		return 3
	}

	if rsArray[K-1][N-1] == 0 {
	loop:
		for floor, floor2 := N/2, N/2; floor > 0; floor, floor2 = floor-1, floor2+1 {
			if N%2 == 0 {
				if rsArray[K-1-1][floor-1-1] == 0 {
					rsArray[K-1-1][floor-1-1] = superEggDrop(K-1, floor-1)
				}

				if rsArray[K-1][floor2-1] == 0 {
					rsArray[K-1][floor2-1] = superEggDrop(K, floor2)
				}

				if rsArray[K-1][floor2-1] == rsArray[K-1-1][floor-1-1] ||
					rsArray[K-1][floor2-1] == rsArray[K-1-1][floor-1-1]+1 {
					rsArray[K-1][N-1] = rsArray[K-1][floor2-1] + 1
					break loop
				}
			} else {
				if rsArray[K-1-1][floor-1] == 0 {
					rsArray[K-1-1][floor-1] = superEggDrop(K-1, floor)
				}

				if rsArray[K-1][floor2-1] == 0 {
					rsArray[K-1][floor2-1] = superEggDrop(K, floor2)
				}

				if rsArray[K-1][floor2-1] == rsArray[K-1-1][floor-1] ||
					rsArray[K-1][floor2-1] == rsArray[K-1-1][floor-1]+1 {
					rsArray[K-1][N-1] = rsArray[K-1][floor2-1] + 1
					break loop
				}
			}
		}
	}

	return rsArray[K-1][N-1]
}

func initSlice(K int, N int) {
	rsArray = make([][]int, K, K)
	for i := 0; i < K; i++ {
		rsArray[i] = make([]int, N, N)
	}
}

func main() {
	fmt.Println(superEggDrop(5, 16))
	fmt.Println(superEggDrop(6, 16))
	fmt.Println(superEggDrop(7, 16))
	fmt.Println(superEggDrop(8, 16))
	fmt.Println(superEggDrop(9, 16))
	fmt.Println(superEggDrop(10, 16))
	fmt.Println(superEggDrop(11, 16))
}
