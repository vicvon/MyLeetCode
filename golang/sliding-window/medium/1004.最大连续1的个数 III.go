/*
 * @lc app=leetcode.cn id=1004 lang=golang
 * @lcpr version=30203
 *
 * [1004] 最大连续1的个数 III
 */

// @lc code=start
func longestOnes(nums []int, k int) int {
	left, right := 0, 0
	curK := k

	maxLen := 0
	for right < len(nums) {
		num := nums[right]
		right++
		if num != 1 {
			curK--
		}

		for curK < 0 && left < right {
			d := nums[left]
			left++
			if d != 1 {
				curK++
			}
		}

		if right-left > maxLen {
			maxLen = right - left
		}
	}

	return maxLen
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,1,0,0,0,1,1,1,1,0]\n2\n
// @lcpr case=end

// @lcpr case=start
// [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1]\n3\n
// @lcpr case=end

*/

