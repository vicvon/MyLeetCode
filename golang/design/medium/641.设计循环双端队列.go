/*
 * @lc app=leetcode.cn id=641 lang=golang
 * @lcpr version=30300
 *
 * [641] 设计循环双端队列
 */

// @lc code=start
type MyCircularDeque struct {
	Cap  int
	Buff []int
	Head int
	Tail int
	Size int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		Cap:  k,
		Buff: make([]int, k),
		Head: 0,
		Tail: k - 1,
		Size: 0,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}

	this.Buff[this.Head] = value
	this.Head = (this.Head + 1) % this.Cap
	this.Size++

	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}

	this.Buff[this.Tail] = value
	if this.Tail == 0 {
		this.Tail = this.Cap - 1
	} else {
		this.Tail = (this.Tail - 1) % this.Cap
	}
	this.Size++

	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}

	if this.Head == 0 {
		this.Head = this.Cap - 1
	} else {
		this.Head--
	}
	this.Size--

	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}

	this.Tail = (this.Tail + 1) % this.Cap
	this.Size--

	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}

	pos := this.Head
	if pos == 0 {
		pos = this.Cap - 1
	} else {
		pos--
	}
	return this.Buff[pos]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.Buff[(this.Tail+1)%this.Cap]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.Size == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.Size == this.Cap
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
// @lc code=end



