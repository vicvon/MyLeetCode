/*
 * @lc app=leetcode.cn id=974 lang=golang
 * @lcpr version=30300
 *
 * [974] 和可被 K 整除的子数组
 */

// @lc code=start
func subarraysDivByK(nums []int, k int) int {
	preSum := make([]int, len(nums)+1)
	preSum[0] = 0
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	valCount := make(map[int]int)
	valCount[0] = 1
	res := 0
	for i := 1; i < len(preSum); i++ {
		remainder := preSum[i] % k
		if remainder < 0 {
			remainder = remainder + k
		}

		if _, ok := valCount[remainder]; !ok {
			valCount[remainder] = 1
		} else {
			res += valCount[remainder]
			valCount[remainder]++
		}
	}

	return res
}

// @lc code=end

/*
// @lcpr case=start
// [4,5,0,-2,-3,1]\n5\n
// @lcpr case=end

// @lcpr case=start
// [5]\n9\n
// @lcpr case=end

*/

