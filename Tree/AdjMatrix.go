package main

import (
	"fmt"
)

const (
	INFINITY int32 = 65536 // 无穷
)

type Edge struct { //边的顶点和权值
	v0     string
	v1     string
	weight int32
}

type Graph struct { //无向网
	v []string //顶点数组
	e []Edge   //边的权值
}

func IndexOfVertex(vs []string, v string) int { //顶点在顶点数组的位置
	for i, value := range vs {
		if value == v {
			return i
		}
	}

	panic("边的顶点不在顶点数组中！")
}

func (this *Graph) AdjMatrix() [][]int32 {
	vertexNum := len(this.v)

	adjM := make([][]int32, vertexNum) //生成矩阵用两次make()
	for i := 0; i < vertexNum; i++ {   //初始化
		adjM[i] = make([]int32, vertexNum)
		for j := 0; j < vertexNum; j++ {
			if j == i {
				adjM[i][j] = 0
			} else {
				adjM[i][j] = INFINITY
			}
		}
	}

	e := this.e
	v := this.v
	for _, edge := range e {
		adjM[IndexOfVertex(v, edge.v0)][IndexOfVertex(v, edge.v1)] = edge.weight
		adjM[IndexOfVertex(v, edge.v1)][IndexOfVertex(v, edge.v0)] = edge.weight //因为是无向图所以是对称矩阵
	}

	return adjM
}

func main() {
	v := []string{"A", "B", "C", "D"}
	e := []Edge{
		Edge{"A", "B", 5},
		Edge{"A", "C", 3},
		Edge{"A", "D", 6},
		Edge{"B", "C", 7},
		Edge{"C", "D", 9},
	}

	graph := &Graph{v, e} //生成无向网

	fmt.Println(graph.AdjMatrix())
	//     A    B     C    D
	//A [  0    5      3    6   ]
	//B [  5    0      7  65536 ]
	//C [  3    7      0    9   ]
	//D [  6  65536    9    0   ]
}
