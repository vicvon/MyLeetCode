/*
 * @lc app=leetcode.cn id=1094 lang=golang
 * @lcpr version=30203
 *
 * [1094] 拼车
 */

// @lc code=start
func carPooling(trips [][]int, capacity int) bool {
	result := make([]int, 1000)
	diff := make([]int, 1000)

	for _, trip := range trips {
		left := trip[1]
		right := trip[2] - 1
		val := trip[0]

		diff[left] += val
		if right+1 < len(diff) {
			diff[right+1] -= val
		}
	}

	result[0] = diff[0]
	for i := 1; i < len(diff); i++ {
		result[i] = diff[i] + result[i-1]
	}

	for _, v := range result {
		if v > capacity {
			return false
		}
	}

	return true
}

// @lc code=end

/*
// @lcpr case=start
// [[2,1,5],[3,3,7]]\n4\n
// @lcpr case=end

// @lcpr case=start
// [[2,1,5],[3,3,7]]\n5\n
// @lcpr case=end

*/

