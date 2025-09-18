/*
 * @lc app=leetcode.cn id=209 lang=golang
 * @lcpr version=30203
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	windowSum := 0
	res := len(nums) + 1
	for right < len(nums) {
		windowSum += nums[right]
		right++

		for windowSum >= target && left < right {
			windowSum -= nums[left]

			if res > right-left {
				res = right - left
			}
			left++
		}
	}

	if res == len(nums)+1 {
		return 0
	}

	return res
}

// @lc code=end

/*
// @lcpr case=start
// 7\n[2,3,1,2,4,3]\n
// @lcpr case=end

// @lcpr case=start
// 4\n[1,4,4]\n
// @lcpr case=end

// @lcpr case=start
// 11\n[1,1,1,1,1,1,1,1]\n
// @lcpr case=end

*/

