/*
 * @lc app=leetcode.cn id=1124 lang=golang
 * @lcpr version=30300
 *
 * [1124] 表现良好的最长时间段
 */

// @lc code=start
func longestWPI(hours []int) int {
	preSum := make([]int, len(hours)+1)
	preSum[0] = 0
	for i := 0; i < len(hours); i++ {
		if hours[i] > 8 {
			preSum[i+1] = preSum[i] + 1
		} else {
			preSum[i+1] = preSum[i] - 1
		}
	}
	valToIdx := make(map[int]int)
	res := 0
	for i := 0; i < len(preSum); i++ {
		if _, ok := valToIdx[preSum[i]]; !ok {
			valToIdx[preSum[i]] = i
		}
		if preSum[i] > 0 {
			res = max(res, i)
		} else {
			if j, found := valToIdx[preSum[i]-1]; found {
				res = max(res, i-j)
			}
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
// [9,9,6,0,6,6,9]\n
// @lcpr case=end

// @lcpr case=start
// [6,6,6]\n
// @lcpr case=end

*/

