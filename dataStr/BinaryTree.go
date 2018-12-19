package dataStr

import (
	"fmt"
)

type BinaryTree struct {
	data  string
	left  *BinaryTree
	right *BinaryTree
}

func PreOrderRec(bt *BinaryTree) { //前序
	if bt == nil {
		return
	}

	fmt.Print(bt.data, " ")
	PreOrderRec(bt.left)
	PreOrderRec(bt.right)
}

func MidOrderRec(bt *BinaryTree) { //中序
	if bt == nil {
		return
	}

	MidOrderRec(bt.left)
	fmt.Print(bt.data, " ")
	MidOrderRec(bt.right)
}

func PostOrderRec(bt *BinaryTree) { //后序
	if bt == nil {
		return
	}

	PostOrderRec(bt.left)
	PostOrderRec(bt.right)
	fmt.Print(bt.data, " ")
}

func PreMid2Tree(pre, mid []string) *BinaryTree { //前序+中序推导二叉树
	if len(pre) != len(mid) {
		panic("两个切片的长度不相等")
	}

	if len(mid) == 0 {
		return nil
	}

	root := &BinaryTree{ //前序第一个元素为root
		data: pre[0],
	}

	if len(mid) == 1 {
		return root
	}

	position := IndexOf(root.data, mid) //找出root在中序的位置

	root.left = PreMid2Tree(pre[1:position+1], mid[:position]) //递归
	root.right = PreMid2Tree(pre[position+1:], mid[position+1:])

	return root
}

func IndexOf(ele string, seq []string) int {
	for i, v := range seq {
		if v == ele {
			return i
		}
	}

	panic("IndexOf错误，元素不存在")
}
