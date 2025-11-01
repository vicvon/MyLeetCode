/*
 * @lc app=leetcode.cn id=560 lang=golang
 * @lcpr version=30300
 *
 * [560] 和为 K 的子数组
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	preSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	valCount := make(map[int]int)
	valCount[0] = 1
	res := 0
	for i := 1; i < len(preSum); i++ {
		target := preSum[i] - k
		if _, ok := valCount[target]; ok {
			res += valCount[target]
		}
		valCount[preSum[i]]++
	}

	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,1]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n3\n
// @lcpr case=end

*/

