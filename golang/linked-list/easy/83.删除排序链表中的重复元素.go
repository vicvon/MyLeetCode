/*
 * @lc app=leetcode.cn id=83 lang=golang
 * @lcpr version=30202
 *
 * [83] 删除排序链表中的重复元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Val: -200}
	dummy.Next = head
	slow, fast := dummy, dummy
	for fast != nil {
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = fast
		}

		fast = fast.Next
	}

	slow.Next = nil
	return dummy.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,2,3,3]\n
// @lcpr case=end

*/

