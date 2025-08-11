/*
 * @lc app=leetcode.cn id=88 lang=golang
 * @lcpr version=30202
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2 := m-1, n-1
	pos := m + n - 1
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[pos] = nums1[p1]
			p1--
		} else {
			nums1[pos] = nums2[p2]
			p2--
		}
		pos--
	}

	for p2 >= 0 {
		nums1[pos] = nums2[p2]
		pos--
		p2--
	}
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,0,0,0]\n3\n[2,5,6]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n[]\n0\n
// @lcpr case=end

// @lcpr case=start
// [0]\n0\n[1]\n1\n
// @lcpr case=end

*/

