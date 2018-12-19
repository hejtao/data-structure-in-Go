package dataStr

var flag []bool //全局变量，用于深度优先遍历

//图
type Edge struct { //给定顶点定义边
	v0 string
	v1 string
}

type Graph struct { //无向图
	v []string //顶点数组
	e []Edge   //边数组
}

//邻接表的数据结构
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

//由图生成邻接表
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

//由图生成邻接矩阵
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

//顶点在顶点数组的位置
func IndexOfVertex(vs []string, v string) int {
	for i, value := range vs {
		if value == v {
			return i
		}
	}

	panic("边的顶点不在顶点数组中！")
}

//由邻接表广度遍历
func BFSTraverse(adjL []LinkList) {
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

//由邻接矩阵深度遍历
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
