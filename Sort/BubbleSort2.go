package main

import (
	"fmt"
)

func BubbleSort2(a []int) {
	flag := true // 有数据交换
	length := len(a)

	for i := 1; i < length-1 && flag; i++ {
		flag = false

		for j := 1; j < length-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = true
			}
		}
	}
}

func main() {
	a := []int{0, 9, 1, 5, 8, 3, 7, 4, 6, 2}
	BubbleSort2(a)
	fmt.Println(a)
}
