#第3章——线性表
- 3.4 顺序存储结构
数值data(起始位置)；数组长度MaxSize；线性表长度length
LOC($a_{i}$)=LOC($a_{1}$)+(i-1)c

- 3.6 链式存储结构
**单链表**，每个节点只包含一个指针域
头指针，头节点，第一个节点

- 3.11 单链表结构与顺序存储结构
内存分配；时间复杂度(查找，插入和删除)；空间复杂度

- 3.12 静态链表

- 3.13 循环链表
尾指针rear
头节点rear->next
第一个节点rear->next->next
合并：
p=rearA->next
q=rearB->next
rearA->next=rearB->next->next
rearB->next=p
free(q)

- 3.14 双向链表
```go
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
	fmt.Print(bl.rear.data)  //打印末节点
}
> Output:
command-line-arguments
  1  2  2.5  3
```
使用标准库
```go
package main

import (
	"container/list"
	"fmt"
)

func main() {
	bl := list.New()
	for i := 1; i < 4; i++ {
		bl.PushBack(i)
	}

	head := bl.Front()
	rear := bl.Back()
	for p := head; p != rear; p = p.Next() {
		fmt.Print(p.Value, "  ")
	}
	fmt.Print(rear.Value)
}
> Output:
command-line-arguments
1  2  3
```


- 4.2 栈
123依次进栈，出栈次序不可能有312(12同时在栈中,2一定先出)
s->top 栈顶指针
栈顶指针为-1表示控栈
s->data[s->top]栈顶元素
```go
type Stack struct { //用于存放 int 的栈
	nums []int
}

func (this *Stack) Push(n int) {
	this.nums = append(this.nums, n)
}

func (this *Stack) Pop() int {
	res := this.nums[len(this.nums)-1]
	this.nums = this.nums[:len(this.nums)-1]
	return res
}
```

- 4.10 队列
```go
type Queue struct { //Queue 是用于存放 int 的队列
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
```

4.12 循环队列
队列满的条件: **(rear+1)%QueueSize == front**
队列长度:**(rear-front+QueueSize)%QueueSize**

- 6.4 树的存储结构
双亲表示(数组)，孩子兄弟表示(数组)，孩子表示(数组+链表)

- 6.5.2 特殊二叉树
斜树，满二叉树，完全二叉树

- 6.6 二叉树性质
总结点数：$n=n_{0}+n_{1}+n_{2}$
分支线总数： $n-1=n_{1}+2n_{2}$
完全二叉树深度：$[log_{2}n]+1$
完全二叉树按层序排号：节点$i$的左节点为$2i$

- 6.8 遍历二叉树
```go
package main

import (
	"fmt"
)

type BinaryTree struct {
	data  string
	left  *BinaryTree
	right *BinaryTree
}

func PreOrderRec(bt *BinaryTree) { //前序
	if bt == nil {
		return
	}

	fmt.Print(bt.data, " ")
	PreOrderRec(bt.left)
	PreOrderRec(bt.right)
}

func MidOrderRec(bt *BinaryTree) { //中序
	if bt == nil {
		return
	}

	MidOrderRec(bt.left)
	fmt.Print(bt.data, " ")
	MidOrderRec(bt.right)
}

func PostOrderRec(bt *BinaryTree) { //后序
	if bt == nil {
		return
	}

	PostOrderRec(bt.left)
	PostOrderRec(bt.right)
	fmt.Print(bt.data, " ")
}

func main() {
	node9 := &BinaryTree{data: "I"}
	node8 := &BinaryTree{data: "H"}
	node7 := &BinaryTree{data: "G"}
	node6 := &BinaryTree{data: "F"}
	node5 := &BinaryTree{data: "E", right: node9}
	node4 := &BinaryTree{"D", node7, node8}
	node3 := &BinaryTree{"C", node5, node6}
	node2 := &BinaryTree{data: "B", left: node4}
	root := &BinaryTree{"A", node2, node3}

	PreOrderRec(root)
	fmt.Println()
	MidOrderRec(root)
	fmt.Println()
	PostOrderRec(root)
}
> Output:
command-line-arguments
A B D G H C E I F 
G D H B A E I C F 
G H D B I E F C A 
```
- 6.8.6 推导遍历结果
前序遍历序列+中序遍历序列->二叉树
后序遍历序列+中序遍历序列->二叉树
前中后：BAC ABC ACB
```go
package main

import (
	"fmt"
)

type BinaryTree struct {
	data  string
	left  *BinaryTree
	right *BinaryTree
}

func PreOrderRec(bt *BinaryTree) { //前序
	if bt == nil {
		return
	}

	fmt.Print(bt.data, " ")
	PreOrderRec(bt.left)
	PreOrderRec(bt.right)
}

func MidOrderRec(bt *BinaryTree) { //中序
	if bt == nil {
		return
	}

	MidOrderRec(bt.left)
	fmt.Print(bt.data, " ")
	MidOrderRec(bt.right)
}

func PostOrderRec(bt *BinaryTree) { //后序
	if bt == nil {
		return
	}

	PostOrderRec(bt.left)
	PostOrderRec(bt.right)
	fmt.Print(bt.data, " ")
}

func PreMid2Tree(pre, mid []string) *BinaryTree { //前序+中序推导二叉树
	if len(pre) != len(mid) {
		panic("两个切片的长度不相等")
	}

	if len(mid) == 0 {
		return nil
	}

	root := &BinaryTree{ //前序第一个元素为root
		data: pre[0],
	}

	if len(mid) == 1 {
		return root
	}

	position := IndexOf(root.data, mid) //找出root在中序的位置

	root.left = PreMid2Tree(pre[1:position+1], mid[:position]) //递归
	root.right = PreMid2Tree(pre[position+1:], mid[position+1:])

	return root
}

func IndexOf(ele string, seq []string) int {
	for i, v := range seq {
		if v == ele {
			return i
		}
	}

	panic("IndexOf错误，元素不存在")
}

func main() {
	bt := PreMid2Tree([]string{"A", "B", "D", "G", "H", "C", "E", "I", "F"},
		[]string{"G", "D", "H", "B", "A", "E", "I", "C", "F"})

	PostOrderRec(bt)
}
> Output:
command-line-arguments
G H D B I E F C A 
```
- 6.10 线索二叉树
将节点的空指针改为指向在遍历序列中的前驱或后继的指针。
```go
type ThreadBiTree struct {
    data  string
    left  *ThreadBiTree
    right *ThreadBiTree
    lTag  *ThreadBiTree //一个bit位，区分是指向孩子还是线索
    rTag  *ThreadBiTree
}

var pre *ThreadBiTree //保存前驱

func MidThreading(bt *ThreadBiTree) { //中序遍历线索化
    if bt == nil {
        return
    }

    MidThreading(bt.left)

    if bt.left == nil { //若bt有左空指针，则把pre设为bt的前驱，并设置标志位
        bt.left = pre
        bt.lTag = 1
    }

    if bt.right == nil { //若bt有右空指针，则把bt设为pre的后继，并设置标志位
        pre.right = bt
        pre.rTag = 1
    }

    pre = bt

    MidThreading(bt.right)
}
```
- 6.11.1 树转换为二叉树
兄弟连线；只保留第一个孩子的连线。

- 6.12 **赫夫曼树**
带权路径长度(WPL)最小的二叉树。
![Huffman Tree](https://upload-images.jianshu.io/upload_images/1863961-81d44580cf5163ee.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)
WPL(a)= 5*1+15*2+40*3+30*4+10*4
WPL(b)= 5*3+15*3+40*2+30*2+10*2

构造
![](https://upload-images.jianshu.io/upload_images/1863961-1acee79132eab45b.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)
排序：A5, E10, B15, D30, C40

7.4 图的存储结构
7.4.1 **邻接矩阵**(adjacency matrix)
无向图：
![](https://upload-images.jianshu.io/upload_images/1863961-d9ce3ee2d67e574e.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)

有向图：
![](https://upload-images.jianshu.io/upload_images/1863961-d915f439c8959688.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)

有向网：
![](https://upload-images.jianshu.io/upload_images/1863961-f214b0dd61c11ca3.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)

无向网的实现：
```go
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
> Output:
command-line-arguments
[[0 5 3 6] [5 0 7 65536] [3 7 0 9] [6 65536 9 0]]
```
7.4.2 **邻接表**
无向图：
![](https://upload-images.jianshu.io/upload_images/1863961-a0854989cc6149ef.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)
```go
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
> Output:
command-line-arguments
v0 1 2 3 
v1 0 2 
v2 0 1 3 
v3 0 2 
```

有向网：
![](https://upload-images.jianshu.io/upload_images/1863961-95004e37703686d9.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)

7.5.1 **深度优先遍历**(Depth First Search)

![](https://upload-images.jianshu.io/upload_images/1863961-4139cb15eb112b82.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)

```go
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
> Output:
command-line-arguments
[0 1 0 0 0 1 0 0 0]
[1 0 1 0 0 0 1 0 1]
[0 1 0 1 0 0 0 0 1]
[0 0 1 0 1 0 1 0 0]
[0 0 0 1 0 1 0 1 0]
[1 0 0 0 1 0 1 0 0]
[0 1 0 1 0 1 0 1 0]
[0 0 0 0 1 0 1 0 0]
[0 1 1 0 0 0 0 0 0]
-------------------
0  1  2  3  4  5  6  7  8  
```
7.5.2 **广度优先遍历**(Breadth First Search)
![](https://upload-images.jianshu.io/upload_images/1863961-1aab42f56392a97a.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)

```go
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
> Output:
command-line-arguments
A B F 
B A C G I 
C B D I 
D C G E 
E F D H 
F A G E 
G B F D H 
H G E 
I B C 
-----------------
A B F C G I E D H 
```
8. 查找
