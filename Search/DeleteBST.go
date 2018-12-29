package main

import (
	"fmt"
)

type BinaryTree struct {
	data  int
	left  *BinaryTree
	right *BinaryTree
}

//删除data为key的节点，并返回该二叉树的根节点。
func DeleteBST(root *BinaryTree, key int) *BinaryTree {
	if root == nil {
		return nil
	}

	switch {
	case key > root.data: //在右子树中
		root.right = DeleteBST(root.right, key)
	case key < root.data: //在左子树中
		root.left = DeleteBST(root.left, key)
	default: //key == root.data
		if root.left == nil && root.right == nil { //该节点为叶节点
			return nil
		} else if root.left == nil && root.right != nil { //该节点仅有右子树
			return root.right
		} else if root.left != nil && root.right == nil { //该节点仅有左子树
			return root.left
		} else { //该节点有左、右子树
			success := FindMin(root.right) //找到key的后继节点,即48
			root.right = DeleteBST(root.right, success)
			root.data = success
		}
	}

	return root
}

//找到BST中data最小的节点
func FindMin(root *BinaryTree) int {
	if root.left == nil { //最小值在根节点
		return root.data
	}

	return FindMin(root.left) //最小值在左子树
}

func main() {
	node16 := &BinaryTree{data: 50}
	node15 := &BinaryTree{data: 48}
	node14 := &BinaryTree{data: 36}
	node13 := &BinaryTree{data: 56}
	node12 := &BinaryTree{49, node15, node16}
	node11 := &BinaryTree{data: 37, left: node14}
	node10 := &BinaryTree{data: 29}
	node9 := &BinaryTree{data: 93}
	node8 := &BinaryTree{51, node12, node13}
	node7 := &BinaryTree{35, node10, node11}
	node6 := &BinaryTree{data: 99, left: node9}
	node5 := &BinaryTree{data: 73}
	node4 := &BinaryTree{data: 47, left: node7, right: node8}
	node3 := &BinaryTree{88, node5, node6}
	node2 := &BinaryTree{data: 58, left: node4}
	root := &BinaryTree{62, node2, node3}

	root = DeleteBST(root, 47)
	NewNode := root.left.left
	fmt.Print(NewNode.data, " ") //打印代替删除位置的新节点
	fmt.Print(NewNode.left.data, NewNode.right.data)
}
