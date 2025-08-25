/*
 * @lc app=leetcode.cn id=707 lang=golang
 * @lcpr version=30202
 *
 * [707] 设计链表
 */

// @lc code=start
type MyLinkedList struct {
	val  int
	next *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{next: nil}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 {
		return -1
	}

	i := 0
	pos := this.next
	for i < index && pos != nil {
		pos = pos.next
		i++
	}

	if pos == nil {
		return -1
	}

	return pos.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	head := &MyLinkedList{val: val, next: this.next}
	this.next = head
}

func (this *MyLinkedList) AddAtTail(val int) {
	tail := &MyLinkedList{val: val, next: nil}
	pre, cur := this, this.next
	for cur != nil {
		pre = cur
		cur = cur.next
	}

	pre.next = tail
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		return
	}

	target := &MyLinkedList{val: val, next: nil}
	i := 0
	pre, cur := this, this.next
	for i < index && cur != nil {
		pre = cur
		cur = cur.next
		i++
	}

	if i < index {
		return
	}

	pre.next = target
	target.next = cur
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 {
		return
	}

	i := 0
	pre, cur := this, this.next
	for i < index && cur != nil {
		pre = cur
		cur = cur.next
		i++
	}

	if i < index || cur == nil {
		return
	}

	pre.next = cur.next
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
// @lc code=end



