/*
 * @lc app=leetcode.cn id=142 lang=golang
 * @lcpr version=30202
 *
 * [142] 环形链表 II
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
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	fast = head
	for fast != nil && slow != nil {
		fast = fast.Next
		slow = slow.Next
		if fast == slow {
			return fast
		}
	}

	return nil
}

// @lc code=end

func TestLinkedListCycleIi(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [3,2,0,-4]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n0\n
// @lcpr case=end

// @lcpr case=start
// [1]\n-1\n
// @lcpr case=end

*/
