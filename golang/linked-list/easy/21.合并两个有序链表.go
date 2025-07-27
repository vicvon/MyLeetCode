/*
 * @lc app=leetcode.cn id=21 lang=golang
 * @lcpr version=30202
 *
 * [21] 合并两个有序链表
 */

package leetcode_solutions

import (
	"fmt"
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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	pos := dummy
	p1 := list1
	p2 := list2
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			pos.Next = p1
			p1 = p1.Next
		} else {
			pos.Next = p2
			p2 = p2.Next
		}
		pos = pos.Next
	}

	if p1 != nil {
		pos.Next = p1
	}
	if p2 != nil {
		pos.Next = p2
	}

	return dummy.Next
}

// @lc code=end

func TestMergeTwoSortedLists(t *testing.T) {
	// your test code here
	list1 := CreateListNode([]int{1, 2, 4})
	list2 := CreateListNode([]int{1, 3, 4})
	result := mergeTwoLists(list1, list2)
	fmt.Println(result)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateListNode(list []int) *ListNode {
	dummy := &ListNode{}
	pos := dummy
	for _, v := range list {
		tmp := &ListNode{Val: v, Next: nil}
		pos.Next = tmp
		pos = pos.Next
	}

	return dummy.Next
}

/*
// @lcpr case=start
// [1,2,4]\n[1,3,4]\n
// @lcpr case=end

// @lcpr case=start
// []\n[]\n
// @lcpr case=end

// @lcpr case=start
// []\n[0]\n
// @lcpr case=end

*/
