/*
 * @lc app=leetcode.cn id=216 lang=golang
 * @lcpr version=30203
 *
 * [216] 组合总和 III
 */

// @lc code=start
func combinationSum3(k int, n int) [][]int {
	var res [][]int
	var track []int
	trackSum := 0

	var backtrack func(int)
	backtrack = func(start int) {
		if trackSum == n && len(track) == k {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}

		if trackSum > n || len(track) > k {
			return
		}

		for i := start; i <= 9; i++ {
			track = append(track, i)
			trackSum += i

			backtrack(i + 1)

			trackSum -= i
			track = track[:len(track)-1]
		}
	}

	backtrack(1)
	return res
}

// @lc code=end

/*
// @lcpr case=start
// 3\n7\n
// @lcpr case=end

// @lcpr case=start
// 3\n9\n
// @lcpr case=end

// @lcpr case=start
// 4\n1\n
// @lcpr case=end

*/

