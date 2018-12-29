package main

import (
	"fmt"
)

type Node struct {
	data string
	pre  *Node
	next *Node
}

type BiLinkList struct {
	length int
	head   *Node
	rear   *Node
}

func NewBiLinkList(head *Node) *BiLinkList {
	return &BiLinkList{0, head, head}
}

func (this *BiLinkList) Append(data string) {
	node := &Node{data: data}

	this.rear.next = node
	node.pre = this.rear
	this.rear = node
	this.length++
}

func (this *BiLinkList) InsertNext(p *Node, e string) {
	//省略判断 nd 不为空且属于链表

	if p.next == nil {
		this.Append(e)
	} else {
		s := &Node{data: e}
		s.pre = p
		s.next = p.next
		p.next.pre = s
		p.next = s
	}
	this.length++
}

func main() {
	head := &Node{}
	bl := NewBiLinkList(head)

	bl.Append("1")
	bl.Append("2")
	bl.Append("3")

	bl.InsertNext(head.next.next, "2.5")

	//for i := 0; i < bl.length; i++ {
	//	fmt.Print(head.next.data, "  ")
	//	head = head.next
	//}

	for node := bl.head; node.next != nil; node = node.next {
		fmt.Print(node.data, "  ")
	}
	fmt.Print(bl.rear.data) //打印末节点
}
