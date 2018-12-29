package main

import (
	"fmt"
)

func BiSearch(a []int, n, key int) int {
	var low, high, mid int
	low = 1
	high = n

	for low <= high {
		mid = (low + high) / 2
		//mid = low + (high-low)(key-1)/(99-1)  插值
		if key == a[mid] {
			return mid
		} else if key < a[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return 0
}

func main() {
	a := []int{0, 1, 16, 24, 35, 47, 59, 62, 73, 88, 99}
	fmt.Println(BiSearch(a, 10, 62))
}
