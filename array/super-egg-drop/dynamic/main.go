package main

import (
	"fmt"
)

var rsArray [][]int

func superEggDrop(K int, N int) int {
	if rsArray == nil {
		initSlice(200, 10000)
	}

	for i := 0; i < K; i++ {
		for j := N - 1; i >= 0; j-- {
			if i == 0 {
				rsArray[i][j] = j
			} else {
				for floor := j; floor > 0{
					if rsArray[i-1][floor] == rsArray[i][N-floor-1]
				}
			}
		}
	}

}

func initSlice(K int, N int) {
	rsArray = make([][]int, K, K)
	for i := 0; i < K; i++ {
		rsArray[i] = make([]int, N, N)
	}
}

func main() {
	fmt.Println(superEggDrop(2, 6))
	fmt.Println(superEggDrop(2, 6))
}
