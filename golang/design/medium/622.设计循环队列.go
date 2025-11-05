/*
 * @lc app=leetcode.cn id=622 lang=golang
 * @lcpr version=30300
 *
 * [622] 设计循环队列
 */

// @lc code=start
type MyCircularQueue struct {
	Cap  int
	Buff []int
	Head int
	Tail int
	Size int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		Cap:  k,
		Buff: make([]int, k),
		Head: 0,
		Tail: 0,
		Size: 0,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	this.Buff[this.Tail] = value
	this.Tail = (this.Tail + 1) % this.Cap
	this.Size++

	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	this.Head = (this.Head + 1) % this.Cap
	this.Size--

	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}

	return this.Buff[this.Head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}

	pos := this.Tail
	if pos == 0 {
		pos = this.Cap - 1
	} else {
		pos--
	}
	return this.Buff[pos]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.Size == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.Size == this.Cap
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
// @lc code=end



