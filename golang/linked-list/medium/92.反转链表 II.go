/*
 * @lc app=leetcode.cn id=92 lang=golang
 * @lcpr version=30202
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var successor *ListNode
	var reverseN func(*ListNode, int) *ListNode
	reverseN = func(head *ListNode, n int) *ListNode {
		if n == 1 {
			successor = head.Next
			return head
		}

		newHead := reverseN(head.Next, n-1)
		head.Next.Next = head
		head.Next = successor
		return newHead
	}

	if left == 1 {
		return reverseN(head, right)
	}

	newHead := reverseBetween(head.Next, left-1, right-1)
	head.Next = newHead
	return head
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n4\n
// @lcpr case=end

// @lcpr case=start
// [5]\n1\n1\n
// @lcpr case=end

*/

