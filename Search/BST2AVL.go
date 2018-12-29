package main

import (
	"fmt"
)

type BinaryTree struct {
	data  int
	bf    int //Balance Factor
	left  *BinaryTree
	right *BinaryTree
}

//右旋
func R_Rotate(root *BinaryTree) *BinaryTree {
	a := root.left
	b := a.right
	a.right = root
	root.left = b

	return a
}

//左旋
func L_Rotate(root *BinaryTree) *BinaryTree {
	a := root.right
	b := a.left
	a.left = root
	root.right = b

	return a
}

//左右旋转
func LR_Rotate(root *BinaryTree) *BinaryTree {
	root.left = L_Rotate(root.left)
	return R_Rotate(root)
}

//右左旋转
func RL_Rotate(root *BinaryTree) *BinaryTree {
	root.right = R_Rotate(root.right)
	return L_Rotate(root)
}

//将任意二叉树转化成AVL
func Balance(root *BinaryTree) *BinaryTree {
	a := &BinaryTree{left: root} //给二叉树生成一个父母节点
	_, isAVL := Bal(root)        //调整二叉树的子树并判断二叉树是否平衡

	for !isAVL {
		Bal(a)                 //处理a的左子树，即二叉树
		_, isAVL = Bal(a.left) //判断二叉树是否平衡
	}

	return a.left //返回二叉树
}

//调整子树
func Bal(root *BinaryTree) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftHeight, leftIsBalanced := Bal(root.left)
	rightHeight, rightIsBalanced := Bal(root.right)

	if !leftIsBalanced {
		root.left = Rotate(root.left)    //调整左子树
		leftHeight = UpdateBF(root.left) //刷新左子树的BF和高度
	}

	if !rightIsBalanced {
		root.right = Rotate(root.right)
		rightHeight = UpdateBF(root.right)
	}

	root.bf = leftHeight - rightHeight //计算本身的BF

	if Abs(root.bf) <= 1 {
		return Max(leftHeight, rightHeight) + 1, true
	}

	return Max(leftHeight, rightHeight) + 1, false
}

//对不平衡树进行旋转调整
func Rotate(root *BinaryTree) *BinaryTree {
	if root.bf > 0 { //左边太重，需要右旋
		if root.left.bf < 0 {
			return LR_Rotate(root)
		}

		return R_Rotate(root)
	}

	if root.right.bf > 0 {
		return RL_Rotate(root)

	}

	return L_Rotate(root)
}

func UpdateBF(root *BinaryTree) int {
	if root == nil {
		return 0
	}

	leftHeight := UpdateBF(root.left)
	rightHeight := UpdateBF(root.right)
	root.bf = leftHeight - rightHeight

	return Max(leftHeight, rightHeight) + 1
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func Abs(a int) int {
	if a > 0 {
		return a
	}

	return -a
}

func InsertBST(root *BinaryTree, key int) (bool, *BinaryTree) {
	//if SearchBST(root, key) {  //假设不存在
	//  return false, root
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

func PrintBF(root *BinaryTree) {
	if root == nil {
		return
	}

	PrintBF(root.left)
	fmt.Print(root.bf, " ")
	PrintBF(root.right)
}

func main() {
	root := &BinaryTree{data: 1}
	InsertBST(root, 7)
	InsertBST(root, 2)
	InsertBST(root, 4)
	InsertBST(root, 8)
	InsertBST(root, 3)
	InsertBST(root, 10)
	InsertBST(root, 5)
	InsertBST(root, 9)
	InsertBST(root, 6)

	UpdateBF(root)
	PrintBF(root) //二叉树的BF
	fmt.Println()
	PrintBF(Balance(root)) //平衡调整后的二叉树的BF
}
