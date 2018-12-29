package main

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

func main() {
	node9 := &BinaryTree{data: "I"}
	node8 := &BinaryTree{data: "H"}
	node7 := &BinaryTree{data: "G"}
	node6 := &BinaryTree{data: "F"}
	node5 := &BinaryTree{data: "E", right: node9}
	node4 := &BinaryTree{"D", node7, node8}
	node3 := &BinaryTree{"C", node5, node6}
	node2 := &BinaryTree{data: "B", left: node4}
	root := &BinaryTree{"A", node2, node3}

	PreOrderRec(root)
	fmt.Println()
	MidOrderRec(root)
	fmt.Println()
	PostOrderRec(root)
}
