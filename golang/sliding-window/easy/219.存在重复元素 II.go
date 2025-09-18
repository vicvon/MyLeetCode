/*
 * @lc app=leetcode.cn id=219 lang=golang
 * @lcpr version=30203
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	left, right := 0, 0
	window := make(map[int]bool)

	for right < len(nums) {
		num := nums[right]
		if window[num] {
			return true
		}
		window[num] = true
		right++

		for right-left > k {
			delete(window, nums[left])
			left++
		}
	}

	return false
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,1]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,0,1,1]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,1,2,3]\n2\n
// @lcpr case=end

*/

