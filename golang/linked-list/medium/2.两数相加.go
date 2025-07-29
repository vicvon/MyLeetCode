/*
 * @lc app=leetcode.cn id=2 lang=golang
 * @lcpr version=30202
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	carry := 0
	p := dummy

	for l1 != nil || l2 != nil || carry > 0 {
		nums := 0
		if l1 != nil {
			nums += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			nums += l2.Val
			l2 = l2.Next
		}

		nums += carry
		curNums := nums % 10
		carry = nums / 10
		p.Next = &ListNode{Val: curNums, Next: nil}
		p = p.Next
	}

	return dummy.Next
}

// @lc code=end

/*
// @lcpr case=start
// [2,4,3]\n[5,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n[0]\n
// @lcpr case=end

// @lcpr case=start
// [9,9,9,9,9,9,9]\n[9,9,9,9]\n
// @lcpr case=end

*/

