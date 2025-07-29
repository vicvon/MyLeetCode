/*
 * @lc app=leetcode.cn id=445 lang=golang
 * @lcpr version=30202
 *
 * [445] 两数相加 II
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
	nums1, nums2 := []int{}, []int{}
	for l1 != nil {
		nums1 = append(nums1, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		nums2 = append(nums2, l2.Val)
		l2 = l2.Next
	}

	dummy := &ListNode{}
	i, j := len(nums1)-1, len(nums2)-1
	carry := 0
	for i >= 0 || j >= 0 || carry > 0 {
		nums := carry
		if i >= 0 {
			nums += nums1[i]
			i--
		}

		if j >= 0 {
			nums += nums2[j]
			j--
		}

		curNum := nums % 10
		carry = nums / 10

		dummy.Next = &ListNode{Val: curNum, Next: dummy.Next}
	}

	return dummy.Next
}

// @lc code=end

/*
// @lcpr case=start
// [7,2,4,3]\n[5,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [2,4,3]\n[5,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n[0]\n
// @lcpr case=end

*/

