package dataStr

type Stack struct { //用于存放 int 的 栈
	nums []int
}

func NewStack() *Stack {
	return &Stack{nums: []int{}}
}

func (this *Stack) Push(n int) {
	this.nums = append(this.nums, n)
}

func (this *Stack) Pop() int {
	res := this.nums[len(this.nums)-1]
	this.nums = this.nums[:len(this.nums)-1]
	return res
}

func (this *Stack) Len() int {
	return len(this.nums)
}

func (this *Stack) IsEmpty() bool {
	return this.Len() == 0
}
