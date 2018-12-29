package main

import (
	"fmt"
)

func ShellSort(a []int) {
	var i, j int
	length := len(a) - 1 //去掉第0位
	inc := length        //增量

	for inc > 1 {
		inc = inc/3 + 1
		for i = 1 + inc; i <= length; i++ { //第(1+inc)个数到最后一个数
			if a[i] < a[i-inc] {
				a[0] = a[i]
				for j = i - inc; j > 0 && a[j] > a[0]; j -= inc {
					a[j+inc] = a[j]
				}
				a[j+inc] = a[0]
			}
		}
	}

}

func main() {
	a := []int{0, 9, 1, 5, 8, 3, 7, 4, 6, 2}
	ShellSort(a)
	fmt.Println(a)
}
