/*
 * @lc app=leetcode.cn id=234 lang=golang
 * @lcpr version=30202
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	left := head
	var right *ListNode

	res := true
	var traverse func(*ListNode)
	traverse = func(r *ListNode) {
		if r == nil {
			return
		}

		traverse(r.Next)

		if left.Val != r.Val {
			res = false
			return
		}

		left = left.Next
	}

	right = head
	traverse(right)
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n
// @lcpr case=end

*/

