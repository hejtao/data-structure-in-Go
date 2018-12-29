package main

import (
	"fmt"
)

func InsertionSort(a []int) {
	length := len(a)
	var j int

	for i := 2; i < length; i++ { //第二个数到最后一个数
		if a[i] < a[i-1] { //第i个数比前面的数小，需要插入
			a[0] = a[i] //哨兵

			for j = i - 1; a[j] > a[0]; j-- { //j取 i-1 到 1
				a[j+1] = a[j] //将大于第i个数的数后移一位，留出空位
			}
			a[j+1] = a[0] //将第i个数放入空位
		}
	}
}

func main() {
	a := []int{0, 9, 1, 5, 8, 3, 7, 4, 6, 2}
	InsertionSort(a)
	fmt.Println(a)
}
