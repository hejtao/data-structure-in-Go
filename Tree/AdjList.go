package main

import (
	"fmt"
)

type Edge struct { //给定顶点定义边
	v0 string
	v1 string
}

type Graph struct { //无向图
	v []string //顶点数组
	e []Edge   //边数组
}

type Node struct { //邻接表的边表节点
	adjvex int
	next   *Node
}

type Vertex struct { //邻接表的顶点
	data      string
	firstedge *Node
}

type LinkList struct { //邻接表的链表
	head *Vertex
	rear *Node
}

func (this *LinkList) Append(adjvex int) {
	newNode := &Node{adjvex: adjvex}

	if this.rear == nil {
		this.head.firstedge = newNode
	} else {
		this.rear.next = newNode
		this.rear = newNode
	}

	this.rear = newNode
}

func IndexOfVertex(vs []string, v string) int { //顶点在顶点数组的位置
	for i, value := range vs {
		if value == v {
			return i
		}
	}

	panic("边的顶点不在顶点数组中！")
}

func (this *Graph) AdjList() []LinkList {
	vertexNum := len(this.v)
	v := this.v
	e := this.e

	adjL := make([]LinkList, vertexNum) //用一个单向链表的数组来表示邻接表

	//for i := 0; i < vertexNum; i++ { //初始
	//  adjL[i].head.data = v[i]                因为此时head为nil，没有data字段
	//}

	for i := 0; i < vertexNum; i++ { //初始
		adjL[i].head = &Vertex{data: v[i]}
	}

	for _, edge := range e {
		i := IndexOfVertex(v, edge.v0)
		j := IndexOfVertex(v, edge.v1)
		adjL[i].Append(j)
		adjL[j].Append(i)
	}

	return adjL
}

func main() {
	v := []string{"v0", "v1", "v2", "v3"}
	e := []Edge{
		Edge{"v0", "v1"},
		Edge{"v0", "v2"},
		Edge{"v0", "v3"},
		Edge{"v1", "v2"},
		Edge{"v2", "v3"},
	}

	graph := &Graph{v, e} //生成无向图

	adjL := graph.AdjList()
	for _, v := range adjL {
		fmt.Print(v.head.data, " ")
		for node := v.head.firstedge; node != nil; node = node.next {
			fmt.Print(node.adjvex, " ")
		}
		fmt.Println()
	}
}
