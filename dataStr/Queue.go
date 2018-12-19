package dataStr

type Queue struct { //Queue 是用于存放 int 的队列
	nums []int
}

func NewQueue() *Queue {
	return &Queue{nums: []int{}}
}

func (this *Queue) Push(n int) {
	this.nums = append(this.nums, n)
}

func (this *Queue) Pop() int {
	res := this.nums[0]
	this.nums = this.nums[1:]
	return res
}

func (this *Queue) Len() int {
	return len(q.nums)
}

func (this *Queue) IsEmpty() bool {
	return this.Len() == 0
}
