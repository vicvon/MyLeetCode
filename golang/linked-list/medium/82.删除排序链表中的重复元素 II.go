/*
 * @lc app=leetcode.cn id=82 lang=golang
 * @lcpr version=30202
 *
 * [82] 删除排序链表中的重复元素 II
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
	dummy := &ListNode{}
	p, q := dummy, head

	for q != nil {
		if q.Next != nil && q.Val == q.Next.Val {
			for q.Next != nil && q.Val == q.Next.Val {
				q = q.Next
			}

			q = q.Next
		} else {
			p.Next = q
			p = p.Next
			q = q.Next
		}
	}

	p.Next = nil

	return dummy.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,3,4,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1,2,3]\n
// @lcpr case=end

*/

