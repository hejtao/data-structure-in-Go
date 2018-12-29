package main

import (
	"fmt"
)

type Node struct {
	data string
	next *Node
}

type LinkList struct {
	length int
	head   *Node
	rear   *Node
}

func NewLinkList(head *Node) *LinkList {
	return &LinkList{0, head, head}
}

func (this *LinkList) Append(data string) {
	if this.rear == nil {
		return
	}
	node := &Node{data: data}

	this.rear.next = node
	this.rear = node
	this.length++
}

func (this *LinkList) Reverse() *LinkList {
	head := this.head
	if head == nil || head.next == nil {
		return this
	}

	var pre *Node = nil
	cur := head.next      //head不为空时，当前为第1节点
	this.rear = head.next //第1节点不为空时，作为最后节点

	for cur != nil {
		cur.next, pre, cur = pre, cur, cur.next

		//buf := cur.next
		//cur.next = pre
		//pre = cur
		//cur = buf
	}
	head.next = pre //指向第最后一个节点

	return this
}

func main() {
	head := &Node{}
	bl := NewLinkList(head)

	bl.Append("1")
	bl.Append("2")
	bl.Append("3")
	bl.Append("4")
	bl.Reverse()

	for node := bl.head; node != nil; node = node.next {
		fmt.Print(node.data, "  ")
	}
}
