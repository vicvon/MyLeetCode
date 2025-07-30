/*
 * @lc app=leetcode.cn id=25 lang=golang
 * @lcpr version=30202
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
var successor *ListNode

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	p := head
	q := head
	for i := 0; i < k; i++ {
		if q == nil {
			return head
		}

		q = q.Next
	}

	newHead := reverseN(p, k)
	head.Next = reverseKGroup(q, k)

	return newHead
}

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		successor = head.Next
		return head
	}

	newHead := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor
	return newHead
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n3\n
// @lcpr case=end

*/

