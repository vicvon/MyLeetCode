/*
 * @lc app=leetcode.cn id=1658 lang=golang
 * @lcpr version=30203
 *
 * [1658] 将 x 减到 0 的最小操作数
 */

// @lc code=start
func minOperations(nums []int, x int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	target := sum - x
	left, right := 0, 0
	windowSum := 0
	maxLen := -1
	for right < len(nums) {
		num := nums[right]
		right++
		windowSum += num

		for windowSum > target && left < right {
			windowSum -= nums[left]
			left++
		}

		if windowSum == target {
			if maxLen == -1 || right-left > maxLen {
				maxLen = right - left
			}
		}
	}

	if maxLen == -1 {
		return -1
	}

	return len(nums) - maxLen
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,4,2,3]\n5\n
// @lcpr case=end

// @lcpr case=start
// [5,6,7,8,9]\n4\n
// @lcpr case=end

// @lcpr case=start
// [3,2,20,1,1,3]\n10\n
// @lcpr case=end

*/

