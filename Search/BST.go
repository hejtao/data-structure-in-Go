package main

import (
	"fmt"
)

type BinaryTree struct {
	data  int
	left  *BinaryTree
	right *BinaryTree
}

func SearchBST(root *BinaryTree, key int) bool {
	if root == nil {
		return false
	}

	switch {
	case key < root.data:
		return SearchBST(root.left, key)
	case key > root.data:
		return SearchBST(root.right, key)
	default:
		return true
	}
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

	fmt.Print(SearchBST(root, 73), " ")
	fmt.Print(SearchBST(root, 99), " ")
	fmt.Print(SearchBST(root, 37), " ")
	fmt.Print(SearchBST(root, 48), " ")
	fmt.Print(SearchBST(root, 100))
}
