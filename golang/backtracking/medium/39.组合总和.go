/*
 * @lc app=leetcode.cn id=39 lang=golang
 * @lcpr version=30203
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var track []int
	trackSum := 0

	var backtrack func(int)
	backtrack = func(start int) {
		if trackSum == target {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}

		if trackSum > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			track = append(track, candidates[i])
			trackSum += candidates[i]

			backtrack(i)

			trackSum -= candidates[i]
			track = track[:len(track)-1]
		}
	}

	backtrack(0)
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [2,3,6,7]\n7\n
// @lcpr case=end

// @lcpr case=start
// [2,3,5]\n8\n
// @lcpr case=end

// @lcpr case=start
// [2]\n1\n
// @lcpr case=end

*/

