
#《大话数据结构》Go 语言

### 3.4 顺序存储结构
数值data(起始位置)；数组长度MaxSize；线性表长度length
LOC($a_{i}$)=LOC($a_{1}$)+(i-1)c

### 3.6 链式存储结构
**单链表**，每个节点只包含一个指针域
头指针，头节点，第一个节点
```go
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
	cur := head.next //head不为空时，当前为第1节点
	this.rear = head.next  //第1节点不为空时，作为最后节点

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
> Output:
command-line-arguments
  4  3  2  1  
```

### 3.11 单链表结构与顺序存储结构
内存分配；时间复杂度(查找，插入和删除)；空间复杂度

### 3.12 静态链表

### 3.13 循环链表
尾指针rear
头节点rear->next
第一个节点rear->next->next
合并：
p=rearA->next
q=rearB->next
rearA->next=rearB->next->next
rearB->next=p
free(q)

### 3.14 双向链表
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


### 4.2 栈
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

### 4.10 队列
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

### 4.12 循环队列
队列满的条件: **(rear+1)%QueueSize == front**
队列长度:**(rear-front+QueueSize)%QueueSize**

### 6.4 树的存储结构
双亲表示(数组)，孩子兄弟表示(数组)，孩子表示(数组+链表)

### 6.5.2 特殊二叉树
斜树，满二叉树，完全二叉树

### 6.6 二叉树性质
总结点数：$n=n_{0}+n_{1}+n_{2}$
分支线总数： $n-1=n_{1}+2n_{2}$
完全二叉树深度：$[log_{2}n]+1$
完全二叉树按层序排号：节点$i$的左节点为$2i$

### 6.8 遍历二叉树
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
### 6.8.6 推导遍历结果
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
### 6.10 线索二叉树
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
### 6.11.1 树转换为二叉树
兄弟连线；只保留第一个孩子的连线

### 6.12 **赫夫曼树**
带权路径长度(WPL)最小的二叉树
<center>![Huffman Tree](https://upload-images.jianshu.io/upload_images/1863961-81d44580cf5163ee.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)

WPL(a)= 5*1+15*2+40*3+30*4+10*4
WPL(b)= 5*3+15*3+40*2+30*2+10*2

构造
<center>![](https://upload-images.jianshu.io/upload_images/1863961-1acee79132eab45b.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)

排序：A5, E10, B15, D30, C40

### 7.4.1 **邻接矩阵**(adjacency matrix)
无向图：
<center>![](https://upload-images.jianshu.io/upload_images/1863961-d9ce3ee2d67e574e.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)

有向图：
<center>![](https://upload-images.jianshu.io/upload_images/1863961-d915f439c8959688.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)

有向网：
<center>![](https://upload-images.jianshu.io/upload_images/1863961-f214b0dd61c11ca3.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)

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
### 7.4.2 **邻接表**
无向图：
<center>![](https://upload-images.jianshu.io/upload_images/1863961-a0854989cc6149ef.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)

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
<center>![](https://upload-images.jianshu.io/upload_images/1863961-95004e37703686d9.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)

### 7.5.1 **深度优先遍历**(Depth First Search)

<center>![](https://upload-images.jianshu.io/upload_images/1863961-4139cb15eb112b82.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)

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
### 7.5.2 **广度优先遍历**(Breadth First Search)
<center>![](https://upload-images.jianshu.io/upload_images/1863961-1aab42f56392a97a.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)


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
### 8.4.1 二分查找
<center>![](https://upload-images.jianshu.io/upload_images/1863961-e7c869cf58b3751b.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)

```go
package main

import (
	"fmt"
)

func BiSearch(a []int, n, key int) int {
	var low, high, mid int
	low = 1
	high = n

	for low <= high {
		mid = (low + high) / 2
		//mid = low + (high-low)(key-1)/(99-1)  插值
		if key == a[mid] {
			return mid
		} else if key < a[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return 0
}

func main() {
	a := []int{0, 1, 16, 24, 35, 47, 59, 62, 73, 88, 99}
	fmt.Println(BiSearch(a, 10, 62))
}
> Output:
command-line-arguments
7
```
### 8.6 二叉查找树(Binary Sort Tree)
左小右大
<center>![](https://upload-images.jianshu.io/upload_images/1863961-ac2cd3676f5e970c.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)
```go
package main

import (
	"fmt"
)

type BinaryTree struct {
	data  int
	left  *BinaryTree
	right *BinaryTree
}

func SearchBST(root *BinaryTree, key int) bool {
	if root == nil {
		return false
	}

	switch {
	case key < root.data:
		return SearchBST(root.left, key)
	case key > root.data:
		return SearchBST(root.right, key)
	default:
		return true
	}
}

func main() {
	node10 := &BinaryTree{data: 37}
	node9 := &BinaryTree{data: 93}
	node8 := &BinaryTree{data: 51}
	node7 := &BinaryTree{data: 35, right: node10}
	node6 := &BinaryTree{data: 99, left: node9}
	node5 := &BinaryTree{data: 73}
	node4 := &BinaryTree{data: 47, left: node7, right: node8}
	node3 := &BinaryTree{88, node5, node6}
	node2 := &BinaryTree{data: 58, left: node4}
	root := &BinaryTree{62, node2, node3}

	fmt.Print(SearchBST(root, 73), " ")
	fmt.Print(SearchBST(root, 99), " ")
	fmt.Print(SearchBST(root, 37), " ")
	fmt.Print(SearchBST(root, 48), " ")
	fmt.Print(SearchBST(root, 100))
}
> Output:
command-line-arguments
true true true false false
```
### 8.6.2 二叉查找树插入
<center>![](https://upload-images.jianshu.io/upload_images/1863961-a95e74e49282a74c.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)

```go
package main

import (
	"fmt"
)

type BinaryTree struct {
	data  int
	left  *BinaryTree
	right *BinaryTree
}

func InsertBST(root *BinaryTree, key int) (bool, *BinaryTree) {
	//if SearchBST(root, key) {  //假设不存在
	//	return false, root
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

func main() {
	node10 := &BinaryTree{data: 37}
	node9 := &BinaryTree{data: 93}
	node8 := &BinaryTree{data: 51}
	node7 := &BinaryTree{data: 35, right: node10}
	node6 := &BinaryTree{data: 99, left: node9}
	node5 := &BinaryTree{data: 73}
	node4 := &BinaryTree{data: 47, left: node7, right: node8}
	node3 := &BinaryTree{88, node5, node6}
	node2 := &BinaryTree{data: 58, left: node4}
	root := &BinaryTree{62, node2, node3}

	fmt.Print(root.right.right.left.data, root.right.right.left.right, " ")
	_, root = InsertBST(root, 95)
	fmt.Print(root.right.right.left.right.data)
}
> Output:
command-line-arguments
93 <nil> 95
```
### 8.6.3 二叉查找树删除
<center>![](https://upload-images.jianshu.io/upload_images/1863961-73f49cf999259ca8.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)

<center>![](https://upload-images.jianshu.io/upload_images/1863961-f901a4ee0de2558c.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)

```go
package main

import (
	"fmt"
)

type BinaryTree struct {
	data  int
	left  *BinaryTree
	right *BinaryTree
}

//删除data为key的节点，并返回该二叉树的根节点。
func DeleteBST(root *BinaryTree, key int) *BinaryTree {
	if root == nil {
		return nil
	}

	switch {
	case key > root.data: //在右子树中
		root.right = DeleteBST(root.right, key)
	case key < root.data: //在左子树中
		root.left = DeleteBST(root.left, key)
	default: //key == root.data
		if root.left == nil && root.right == nil { //该节点为叶节点
			return nil
		} else if root.left == nil && root.right != nil { //该节点仅有右子树
			return root.right
		} else if root.left != nil && root.right == nil { //该节点仅有左子树
			return root.left
		} else { //该节点有左、右子树
			success := FindMin(root.right) //找到key的后继节点,即48
			root.right = DeleteBST(root.right, success)
			root.data = success
		}
	}

	return root
}

//找到BST中data最小的节点
func FindMin(root *BinaryTree) int {
	if root.left == nil { //最小值在根节点
		return root.data
	}

	return FindMin(root.left) //最小值在左子树
}

func main() {
	node16 := &BinaryTree{data: 50}
	node15 := &BinaryTree{data: 48}
	node14 := &BinaryTree{data: 36}
	node13 := &BinaryTree{data: 56}
	node12 := &BinaryTree{49, node15, node16}
	node11 := &BinaryTree{data: 37, left: node14}
	node10 := &BinaryTree{data: 29}
	node9 := &BinaryTree{data: 93}
	node8 := &BinaryTree{51, node12, node13}
	node7 := &BinaryTree{35, node10, node11}
	node6 := &BinaryTree{data: 99, left: node9}
	node5 := &BinaryTree{data: 73}
	node4 := &BinaryTree{data: 47, left: node7, right: node8}
	node3 := &BinaryTree{88, node5, node6}
	node2 := &BinaryTree{data: 58, left: node4}
	root := &BinaryTree{62, node2, node3}

	root = DeleteBST(root, 47)
	NewNode := root.left.left
	fmt.Print(NewNode.data, " ") //打印代替删除位置的新节点
	fmt.Print(NewNode.left.data, NewNode.right.data)
}
> Output:
command-line-arguments
48 35 51
```
### 8.7 平衡二叉树(AVL树)
**平衡因子**(BF)=左子树的深度-右子树的深度
**旋转：**
<center>![](https://upload-images.jianshu.io/upload_images/1863961-b493bc28e8068b46.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)
<center>![](https://upload-images.jianshu.io/upload_images/1863961-b6715f1391c81757.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)
```go
//右旋
func RRotate(k2 *BinaryTree) *BinaryTree {
	k1 := k2.left
	y := k1.right
	k1.right = k2
	k2.left = y
	
	return k1
}
```
**双旋转：**
<center>![](https://upload-images.jianshu.io/upload_images/1863961-132deac059f8995f.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)
```go
//左右旋转
func LRRotate(k3 *BinaryTree) *BinaryTree {
	k3.left = LRotate(k3.left)
	return RRotate(k3)
}
```
<center>![](https://upload-images.jianshu.io/upload_images/1863961-fe05b14c2bdb3e2c.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/800)

将任意二叉树一次性调整AVL
```go
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
> Output:
command-line-arguments
-5 -3 0 -1 -1 0 1 -2 0 1 
0 0 0 -1 -1 0 0 0 0 0 
```
### 8.9  散列表(哈希表)查找
每一个关键字 $key$ 对应一个存储位置 $f(key)$





### 9.2.1 排序的稳定性
<center>![](https://upload-images.jianshu.io/upload_images/1863961-da3eb9745c415426.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)

### 9.2.1 内排序和外排序
**内排序**是在排序的整个过程中，待排序的所有记录全部在内存中。**外排序**的整个过程则需要在内外存之间交换数据。

### 9.3.2 冒泡排序算法(Bubble Sort)
![](https://upload-images.jianshu.io/upload_images/1863961-35aca3bf05e9be84.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)

```go
package main

import (
	"fmt"
)

package main

import (
	"fmt"
)

func BubbleSort(a []int) {
	length := len(a)
	for i := 1; i < length-1; i++ { //需要交换(length-2)次
		for j := 1; j < length-i; j++ { //当i=1时，j可以取到(length-2)
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func main() {
	a := []int{0, 9, 1, 5, 8, 3, 7, 4, 6, 2}
	BubbleSort(a)
	fmt.Println(a)
}
```
### 9.3.3 冒泡排序优化
<center>![](https://upload-images.jianshu.io/upload_images/1863961-2c7713c7b8115947.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
```go
package main

import (
	"fmt"
)

func BubbleSort2(a []int) {
	flag := true // 有数据交换
	length := len(a)

	for i := 1; i < length-1 && flag; i++ { 
		flag = false

		for j := 1; j < length-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = true
			}
		}
	}
}

func main() {
	a := []int{0, 9, 1, 5, 8, 3, 7, 4, 6, 2}
	BubbleSort2(a)
	fmt.Println(a)
}
```
### 9.4.1 简单选择排序(Simple Selection Sort)
复杂度与冒泡排序同为$O(n^{2})$,但性能更优(数据交换次数更少)
```go
package main

import (
	"fmt"
)

func SelectionSort(a []int) {
	var min int
	length := len(a)

	for i := 1; i < length-1; i++ { //比较次数(length-2)
		min = i

		for j := i + 1; j < length; j++ { //j取到i后的所有数
			if a[j] < a[min] { //之后有更小的数
				j, min = min, j //交换下标
			}
		}

		if i != min { //下标改变，交换
			a[i], a[min] = a[min], a[i]
		}
	}

}

func main() {
	a := []int{0, 9, 1, 5, 8, 3, 7, 4, 6, 2}
	SelectionSort(a)
	fmt.Println(a)
}
```
### 9.5.1 直接插入排序(Straight Insertion Sort)
复杂度同为为$O(n^{2})$，性能：**插入排序>选择排序>冒泡排序**
```go
package main

import (
	"fmt"
)

func InsertionSort(a []int) {
	length := len(a)
	var j int

	for i := 2; i < length; i++ { //第二个数到最后一个数
		if a[i] < a[i-1] { //第i个数比前面的数小，需要插入
			a[0] = a[i] //哨兵

			for j = i - 1; a[j] > a[0]; j-- { //j取 i-1 到 1
				a[j+1] = a[j] //将大于第i个数的数后移一位，留出空位
			}
			a[j+1] = a[0] //将第i个数放入空位
		}
	}
}

func main() {
	a := []int{0, 9, 1, 5, 8, 3, 7, 4, 6, 2}
	InsertionSort(a)
	fmt.Println(a)
}
```
### 9.6 希尔排序(Shell Sort)
按照间隔分组直接插入排序
<center>![](https://upload-images.jianshu.io/upload_images/1863961-8713bb4788430478.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)
```go
package main

import (
	"fmt"
)

func ShellSort(a []int) {
	var i, j int
	length := len(a) - 1 //去掉第0位
	inc := length //增量

	for inc > 1 {
		inc = inc/3 + 1
		for i = 1 + inc; i <= length; i++ { //第(1+inc)个数到最后一个数
			if a[i] < a[i-inc] {
				a[0] = a[i]
				for j = i - inc; j > 0 && a[j] > a[0]; j -= inc {
					a[j+inc] = a[j]
				}
				a[j+inc] = a[0]
			}
		}
	}

}

func main() {
	a := []int{0, 9, 1, 5, 8, 3, 7, 4, 6, 2}
	ShellSort(a)
	fmt.Println(a)
}
```
### 9.7 堆排序
时间复杂度$O(nlogn)$
```go
package main

import (
	"fmt"
)

func HeapSort(a []int) {
	length := len(a) - 1

	//循环后a[1]为最大值
	for i := length / 2; i > 0; i-- { //(length/2)是最后一个节点的父节点到根节点
		HeapAdjust(a, i, length)
	}

	for i := length; i > 1; i-- { //从最后节点到第二个节点
		a[1], a[i] = a[i], a[1] //排序第i位
		HeapAdjust(a, 1, i-1)   //将1到i-1中的最大数放到a[1]
	}
}

func HeapAdjust(a []int, s, m int) {
	var temp, j int
	temp = a[s]

	for j = 2 * s; j <= m; j *= 2 { //以s为父节点开始
		if j < m && a[j] < a[j+1] { //取出较大的孩子节点
			j = j + 1
		}

		if temp >= a[j] { //父节点已经最大
			break
		}
		a[s] = a[j] //将最大的值替换给父节点
		s = j       //将当前节点作为父节点，进行下一轮操作
	}
	a[s] = temp
}

func main() {
	a := []int{0, 9, 1, 5, 8, 3, 7, 4, 6, 2}
	HeapAdjust(a, 1, 9)
	fmt.Println(a)
}
```
### 9.8 归并排序
递归方法
<center>![](https://upload-images.jianshu.io/upload_images/1863961-494a82286018f65f.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)

Merge()归并排序示意图：
<center>![](https://upload-images.jianshu.io/upload_images/1863961-cfe727471bdad794.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)


```go
package main

import (
	"fmt"
)

// 将a排序到b
func MergeSort(a []int) []int {
	length := len(a)
	b := make([]int, length)

	MSort(a, b, 1, length-1)

	return b
}

func MSort(a, b []int, s, t int) {
	if s == t {
		b[s] = a[s] //将a复制到到b
	} else {
		m := (s + t) / 2
		MSort(a, b, s, m)
		MSort(a, b, m+1, t)
		Merge(b, s, m, t) //将b归并排序
	}
}

func Merge(SR []int, i, m, n int) {
	TR := make([]int, len(SR)) //归并的序列暂存到TR

	s := i //保存起始位置
	j := m + 1
	k := i //TR序号

	for i <= m && j <= n {
		if SR[i] < SR[j] {
			TR[k] = SR[i]
			i++
		} else {
			TR[k] = SR[j]
			j++
		}

		k++
	}

	if i <= m {
		for l := 0; l <= m-i; l++ {
			TR[k+l] = SR[i+l]
		}
	}

	if j <= n {
		for l := 0; l <= n-j; l++ {
			TR[k+l] = SR[j+l]
		}
	}

	for p := s; p <= n; p++ { //将排好序的TR写回到SR
		if SR[p] != TR[p] {
			SR[p] = TR[p]
		}
	}
}

func main() {
	a := []int{0, 9, 1, 5, 8, 3, 7, 4, 6, 5, 16}
	fmt.Println(MergeSort(a))
}
```
非递归方法
```go
package main

import (
	"fmt"
)

func MergeSort2(a []int) {
	length := len(a) - 1

	for k := 1; k < length; {
		MergePass(a, k, length)
		k = k * 2
	}
}

func MergePass(a []int, s, n int) {
	i := 1
	for i <= n-2*s+1 {
		Merge2(a, i, i+s-1, i+2*s-1) //i+2*s-1<=n
		i = i + 2*s
	}

	if i < n-s+1 { //n>i+s-1,归并最后两个子块
		Merge2(a, i, i+s-1, n)
	}
}

func Merge2(SR []int, i, m, n int) {
	TR := make([]int, n-i+1) //与Merge相比栈空间更小

	s := i //保存起始位置
	j := m + 1
	k := 0 //TR序号

	for i <= m && j <= n {
		if SR[i] < SR[j] {
			TR[k] = SR[i]
			i++
		} else {
			TR[k] = SR[j]
			j++
		}

		k++
	}

	if i <= m {
		for l := 0; l <= m-i; l++ {
			TR[k+l] = SR[i+l]
		}
	}

	if j <= n {
		for l := 0; l <= n-j; l++ {
			TR[k+l] = SR[j+l]
		}
	}

	for p := s; p <= n; p++ { //将排好序的TR写回到SR
		if SR[p] != TR[p-s] {
			SR[p] = TR[p-s]
		}
	}
}

func main() {
	a := []int{0, 9, 1, 5, 8, 3, 7, 4, 6, 5, 16, 3, 41, 7, 55, 21}
	MergeSort2(a)
	fmt.Println(a)
}
```
### 9.9 快速排序
```go
package main

import (
	"fmt"
)

func QuickSort(a []int) {
	Qsort(a, 1, len(a)-1)
}

func Qsort(a []int, low, high int) {
	if low < high {
		pivot := Partition(a, low, high)

		Qsort(a, low, pivot-1)
		Qsort(a, pivot+1, high)
	}

}

func Partition(a []int, low, high int) int {
	pivotValue := a[low]

	for low < high {
		for pivotValue <= a[high] { //找出a[high]<pivotValue
			high--
		}
		a[low], a[high] = a[high], a[low] //将a[hign]放到pivoValue左边

		for low < high && pivotValue >= a[low] { //找出a[low]>pivotValue
			low++
		}
		a[low], a[high] = a[high], a[low] //将a[low]放到pivoValue右边
	}

	return low
}

func Partition2(a []int, low, high int) int {
	pivotValue := a[low]

	for low < high {
		for pivotValue <= a[high] { //找出a[high]<pivotValue
			high--
		}
		a[low] = a[high] //将a[hign]放到较低的位置

		for low < high && pivotValue >= a[low] { //找出a[low]>pivotValue
			low++
		}
		a[high] = a[low] //将a[low]放到较高的位置
	}
	a[low] = pivotValue
	return low
}

func main() {
	a := []int{0, 50, 10, 90, 30, 70, 40, 80, 60, 20}
	QuickSort(a)
	fmt.Println(a)
}
```
<center>![](https://upload-images.jianshu.io/upload_images/1863961-25127e3a301cf527.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/640)




