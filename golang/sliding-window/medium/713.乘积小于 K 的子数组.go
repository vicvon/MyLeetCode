/*
 * @lc app=leetcode.cn id=713 lang=golang
 * @lcpr version=30203
 *
 * [713] 乘积小于 K 的子数组
 */

// @lc code=start
func numSubarrayProductLessThanK(nums []int, k int) int {
	left, right := 0, 0
	windowProduct := 1
	count := 0

	for right < len(nums) {
		windowProduct *= nums[right]
		right++

		for left < right && windowProduct >= k {
			windowProduct /= nums[left]
			left++
		}

		count += right - left
	}

	return count
}

// @lc code=end

/*
// @lcpr case=start
// [10,5,2,6]\n100\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n0\n
// @lcpr case=end

*/

