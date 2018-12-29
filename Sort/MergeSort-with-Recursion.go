package main

import (
	"fmt"
)

// 将a排序到b
func MergeSort(a []int) []int {
	length := len(a)
	b := make([]int, length)

	MSort(a, b, 1, length-1)

	return b
}

func MSort(a, b []int, s, t int) {
	if s == t {
		b[s] = a[s] //将a复制到到b
	} else {
		m := (s + t) / 2
		MSort(a, b, s, m)
		MSort(a, b, m+1, t)
		Merge(b, s, m, t) //将b归并排序
	}
}

func Merge(SR []int, i, m, n int) {
	TR := make([]int, len(SR)) //归并的序列暂存到TR

	s := i //保存起始位置
	j := m + 1
	k := i //TR序号

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
		if SR[p] != TR[p] {
			SR[p] = TR[p]
		}
	}
}

func main() {
	a := []int{0, 9, 1, 5, 8, 3, 7, 4, 6, 5, 16}
	fmt.Println(MergeSort(a))
}
