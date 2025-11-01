/*
 * @lc app=leetcode.cn id=523 lang=golang
 * @lcpr version=30300
 *
 * [523] 连续的子数组和
 */

// @lc code=start
func checkSubarraySum(nums []int, k int) bool {
	preSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	valToIdx := make(map[int]int)
	for i := 0; i < len(preSum); i++ {
		val := preSum[i] % k
		if _, ok := valToIdx[val]; !ok {
			valToIdx[val] = i
		} else {
			if i-valToIdx[val] >= 2 {
				return true
			}
		}
	}

	return false
}

// @lc code=end

/*
// @lcpr case=start
// [23,2,4,6,7]\n6\n
// @lcpr case=end

// @lcpr case=start
// [23,2,6,4,7]\n6\n
// @lcpr case=end

// @lcpr case=start
// [23,2,6,4,7]\n13\n
// @lcpr case=end

*/

