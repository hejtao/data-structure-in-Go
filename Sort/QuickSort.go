package main

import (
	"fmt"
)

func QuickSort(a []int) {
	Qsort(a, 1, len(a)-1)
}

func Qsort(a []int, low, high int) {
	if low < high {
		pivot := Partition(a, low, high)

		Qsort(a, low, pivot-1)
		Qsort(a, pivot+1, high)
	}

}

func Partition(a []int, low, high int) int {
	pivotValue := a[low]

	for low < high {
		for pivotValue <= a[high] { //找出a[high]<pivotValue
			high--
		}
		a[low], a[high] = a[high], a[low] //将a[hign]放到pivoValue左边

		for low < high && pivotValue >= a[low] { //找出a[low]>pivotValue
			low++
		}
		a[low], a[high] = a[high], a[low] //将a[low]放到pivoValue右边
	}

	return low
}

func Partition2(a []int, low, high int) int {
	pivotValue := a[low]

	for low < high {
		for pivotValue <= a[high] { //找出a[high]<pivotValue
			high--
		}
		a[low] = a[high] //将a[hign]放到较低的位置

		for low < high && pivotValue >= a[low] { //找出a[low]>pivotValue
			low++
		}
		a[high] = a[low] //将a[low]放到较高的位置
	}
	a[low] = pivotValue
	return low
}

func main() {
	a := []int{0, 50, 10, 90, 30, 70, 40, 80, 60, 20}
	QuickSort(a)
	fmt.Println(a)
}
