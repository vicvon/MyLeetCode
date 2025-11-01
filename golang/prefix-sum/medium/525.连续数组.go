/*
 * @lc app=leetcode.cn id=525 lang=golang
 * @lcpr version=30300
 *
 * [525] 连续数组
 */

// @lc code=start
func findMaxLength(nums []int) int {
	preSum := make([]int, len(nums)+1)
	preSum[0] = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			preSum[i+1] = preSum[i] - 1
		} else {
			preSum[i+1] = preSum[i] + 1
		}
	}

	valToIdx := make(map[int]int)
	res := 0
	for i := 0; i < len(preSum); i++ {
		if _, ok := valToIdx[preSum[i]]; !ok {
			valToIdx[preSum[i]] = i
		} else {
			res = max(res, i-valToIdx[preSum[i]])
		}
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

// @lc code=end

/*
// @lcpr case=start
// [0,1]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,0]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,1,1,1,1,0,0,0]\n
// @lcpr case=end

*/

