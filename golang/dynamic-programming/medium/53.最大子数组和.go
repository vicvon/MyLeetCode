/*
 * @lc app=leetcode.cn id=53 lang=golang
 * @lcpr version=30203
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := dp[0]

	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}

		if res < dp[i] {
			res = dp[i]
		}
	}

	return res
}

// @lc code=end

/*
// @lcpr case=start
// [-2,1,-3,4,-1,2,1,-5,4]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// [5,4,-1,7,8]\n
// @lcpr case=end

*/

