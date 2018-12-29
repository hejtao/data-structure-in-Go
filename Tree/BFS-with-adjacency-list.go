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

type Queue struct {
	nums []int
}

func (this *Queue) Push(n int) {
	this.nums = append(this.nums, n)
}

func (this *Queue) Pop() int {
	res := this.nums[0]
	this.nums = this.nums[1:]
	return res
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
	//	adjL[i].head.data = v[i]                因为此时head为nil，没有data字段
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

func BFSTraverse(adjL []LinkList) { //输入邻接表
	vertexNum := len(adjL)

	flag := make([]bool, vertexNum) //标记被访问的节点
	for i := 0; i < vertexNum; i++ {
		flag[i] = false
	}

	Q := new(Queue) //队列

	for i := 0; i < vertexNum; i++ {
		if !flag[i] { //如果顶点未被访问
			flag[i] = true
			Q.Push(i)

			for len(Q.nums) > 0 { //队列不为空
				j := Q.Pop()
				fmt.Print(adjL[j].head.data, " ") //打印节点
				for node := adjL[j].head.firstedge; node != nil; node = node.next {
					if flag[node.adjvex] == false {
						flag[node.adjvex] = true
						Q.Push(node.adjvex)
					}
				}

			}

		}

	}
}

func main() {
	v := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	e := []Edge{
		Edge{"A", "B"},
		Edge{"A", "F"},
		Edge{"B", "C"},
		Edge{"B", "G"},
		Edge{"B", "I"},
		Edge{"F", "G"},
		Edge{"F", "E"},
		Edge{"C", "D"},
		Edge{"C", "I"},
		Edge{"G", "D"},
		Edge{"G", "H"},
		Edge{"E", "D"},
		Edge{"E", "H"},
	}

	graph := &Graph{v, e} //生成无向图

	adjL := graph.AdjList()
	for _, v := range adjL {
		fmt.Print(v.head.data, " ")
		for node := v.head.firstedge; node != nil; node = node.next {
			fmt.Print(adjL[node.adjvex].head.data, " ")
		}
		fmt.Println()
	}
	fmt.Println("-----------------")
	BFSTraverse(adjL)
}
