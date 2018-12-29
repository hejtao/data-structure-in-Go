package main

import (
	"fmt"
)

func HeapSort(a []int) {
	length := len(a) - 1

	//循环后a[1]为最大值
	for i := length / 2; i > 0; i-- { //(length/2)是最后一个节点的父节点到根节点
		HeapAdjust(a, i, length)
	}

	for i := length; i > 1; i-- { //从最后节点到第二个节点
		a[1], a[i] = a[i], a[1] //排序第i位
		HeapAdjust(a, 1, i-1)   //将1到i-1中的最大数放到a[1]
	}
}

func HeapAdjust(a []int, s, m int) {
	var temp, j int
	temp = a[s]

	for j = 2 * s; j <= m; j *= 2 { //以s为父节点开始
		if j < m && a[j] < a[j+1] { //取出较大的孩子节点
			j = j + 1
		}

		if temp >= a[j] { //父节点已经最大
			break
		}
		a[s] = a[j] //将最大的值替换给父节点
		s = j       //将当前节点作为父节点，进行下一轮操作
	}
	a[s] = temp
}

func main() {
	a := []int{0, 9, 1, 5, 8, 3, 7, 4, 6, 2}
	HeapAdjust(a, 1, 9)
	fmt.Println(a)
}
