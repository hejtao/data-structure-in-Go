package dataStr

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
