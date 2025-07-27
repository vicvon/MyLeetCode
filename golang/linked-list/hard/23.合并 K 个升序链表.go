/*
 * @lc app=leetcode.cn id=23 lang=golang
 * @lcpr version=30202
 *
 * [23] 合并 K 个升序链表
 */

package leetcode_solutions

import (
	"container/heap"
	"testing"
)

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type MinHeap []*ListNode

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	dummy := &ListNode{}
	p := dummy

	h := &MinHeap{}
	heap.Init(h)

	for _, head := range lists {
		if head != nil {
			heap.Push(h, head)
		}
	}

	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		p.Next = node
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
		p = p.Next
	}

	return dummy.Next
}

// @lc code=end

func TestMergeKSortedLists(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [[1,4,5],[1,3,4],[2,6]]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [[]]\n
// @lcpr case=end

*/
