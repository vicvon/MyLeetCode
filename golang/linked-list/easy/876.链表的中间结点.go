/*
 * @lc app=leetcode.cn id=876 lang=golang
 * @lcpr version=30202
 *
 * [876] 链表的中间结点
 */

package leetcode_solutions

import "testing"

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	fast := dummy
	slow := dummy
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	if fast == nil {
		return slow
	}

	return slow.Next
}

// @lc code=end

func TestMiddleOfTheLinkedList(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5,6]\n
// @lcpr case=end

*/
