package main

import (
	"fmt"
)

var flag []bool //全局变量

type Edge struct { //给定顶点定义边
	v0 string
	v1 string
}

type Graph struct { //无向图
	v []string //顶点数组
	e []Edge   //边数组
}

func (this *Graph) AdjMatrix() [][]int {
	vertexNum := len(this.v)

	adjM := make([][]int, vertexNum) //生成矩阵用两次make()
	for i := 0; i < vertexNum; i++ { //初始化
		adjM[i] = make([]int, vertexNum)
		for j := 0; j < vertexNum; j++ {
			adjM[i][j] = 0
		}
	}

	e := this.e
	v := this.v
	for _, edge := range e {
		adjM[IndexOfVertex(v, edge.v0)][IndexOfVertex(v, edge.v1)] = 1
		adjM[IndexOfVertex(v, edge.v1)][IndexOfVertex(v, edge.v0)] = 1
	}

	return adjM
}

func IndexOfVertex(vs []string, v string) int {
	for i, value := range vs {
		if value == v {
			return i
		}
	}

	panic("边的顶点不在顶点数组中！")
}

func DFSTraverse(adjM [][]int) { //输入无向图的邻接矩阵
	vertexNum := len(adjM)

	flag = make([]bool, vertexNum) //初始化flag，注意这里不能用 :=
	for i := 0; i < vertexNum; i++ {
		flag[i] = false
	}

	for i := 0; i < vertexNum; i++ { //对未访问的节点执行深度搜索
		if !flag[i] {
			DFS(vertexNum, i, adjM)
		}
	}
}

func DFS(vertexNum int, i int, adjM [][]int) {
	flag[i] = true
	fmt.Print(i, "  ")               //打印被访问的节点序号
	for j := 0; j < vertexNum; j++ { //对未被访问的邻节点递归
		if adjM[i][j] == 1 && !flag[j] {
			DFS(vertexNum, j, adjM)
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
	adjM := graph.AdjMatrix()
	for i := 0; i < len(adjM); i++ {
		fmt.Println(adjM[i])
	}
	fmt.Println("-------------------")
	DFSTraverse(adjM)
}
