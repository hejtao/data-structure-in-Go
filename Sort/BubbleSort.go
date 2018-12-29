package main

import (
	"fmt"
)

func BubbleSort(a []int) {
	length := len(a)
	for i := 1; i < length-1; i++ { //需要交换(length-2)次
		for j := 1; j < length-i; j++ { //当i=1时，j可以取到(length-2)
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func main() {
	a := []int{0, 9, 1, 5, 8, 3, 7, 4, 6, 2}
	BubbleSort(a)
	fmt.Println(a)
}
