package main

import (
	"fmt"
)

func SelectionSort(a []int) {
	var min int
	length := len(a)

	for i := 1; i < length-1; i++ { //比较次数(length-2)
		min = i

		for j := i + 1; j < length; j++ { //j取到i后的所有数
			if a[j] < a[min] { //之后有更小的数
				j, min = min, j //交换下标
			}
		}

		if i != min { //下标改变，交换
			a[i], a[min] = a[min], a[i]
		}
	}

}

func main() {
	a := []int{0, 9, 1, 5, 8, 3, 7, 4, 6, 2}
	SelectionSort(a)
	fmt.Println(a)
}
