/*
 * @lc app=leetcode.cn id=40 lang=golang
 * @lcpr version=30203
 *
 * [40] 组合总和 II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var track []int
	trackSum := 0

	sort.Ints(candidates)
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
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			track = append(track, candidates[i])
			trackSum += candidates[i]

			backtrack(i + 1)

			track = track[:len(track)-1]
			trackSum -= candidates[i]
		}
	}

	backtrack(0)
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [10,1,2,7,6,1,5]\n8\n
// @lcpr case=end

// @lcpr case=start
// [2,5,2,1,2]\n5\n
// @lcpr case=end

*/

