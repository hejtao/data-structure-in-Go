package main

import (
	"fmt"
)

type BinaryTree struct {
	data  int
	left  *BinaryTree
	right *BinaryTree
}

func InsertBST(root *BinaryTree, key int) (bool, *BinaryTree) {
	//if SearchBST(root, key) {  //假设不存在
	//	return false, root
	//}

	return true, Insert(root, key)
}

func Insert(root *BinaryTree, key int) *BinaryTree {
	if root == nil {
		return &BinaryTree{data: key} //插入的本质要生成新的节点
	}

	if key < root.data {
		root.left = Insert(root.left, key)
	} else { // 没有key = root.data 的情况
		root.right = Insert(root.right, key)
	}

	return root
}

func main() {
	node10 := &BinaryTree{data: 37}
	node9 := &BinaryTree{data: 93}
	node8 := &BinaryTree{data: 51}
	node7 := &BinaryTree{data: 35, right: node10}
	node6 := &BinaryTree{data: 99, left: node9}
	node5 := &BinaryTree{data: 73}
	node4 := &BinaryTree{data: 47, left: node7, right: node8}
	node3 := &BinaryTree{88, node5, node6}
	node2 := &BinaryTree{data: 58, left: node4}
	root := &BinaryTree{62, node2, node3}

	fmt.Print(root.right.right.left.data, root.right.right.left.right, " ")
	_, root = InsertBST(root, 95)
	fmt.Print(root.right.right.left.right.data)
}
