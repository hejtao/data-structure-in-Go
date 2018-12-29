package main

import (
	"fmt"
)

func MergeSort2(a []int) {
	length := len(a) - 1

	for k := 1; k < length; {
		MergePass(a, k, length)
		k = k * 2
	}
}

func MergePass(a []int, s, n int) {
	i := 1
	for i <= n-2*s+1 {
		Merge2(a, i, i+s-1, i+2*s-1) //i+2*s-1<=n
		i = i + 2*s
	}

	if i < n-s+1 { //n>i+s-1,归并最后两个子块
		Merge2(a, i, i+s-1, n)
	}
}

func Merge2(SR []int, i, m, n int) {
	TR := make([]int, n-i+1) //与Merge相比栈空间更小

	s := i //保存起始位置
	j := m + 1
	k := 0 //TR序号

	for i <= m && j <= n {
		if SR[i] < SR[j] {
			TR[k] = SR[i]
			i++
		} else {
			TR[k] = SR[j]
			j++
		}

		k++
	}

	if i <= m {
		for l := 0; l <= m-i; l++ {
			TR[k+l] = SR[i+l]
		}
	}

	if j <= n {
		for l := 0; l <= n-j; l++ {
			TR[k+l] = SR[j+l]
		}
	}

	for p := s; p <= n; p++ { //将排好序的TR写回到SR
		if SR[p] != TR[p-s] {
			SR[p] = TR[p-s]
		}
	}
}

func main() {
	a := []int{0, 9, 1, 5, 8, 3, 7, 4, 6, 5, 16, 3, 41, 7, 55, 21}
	MergeSort2(a)
	fmt.Println(a)
}
