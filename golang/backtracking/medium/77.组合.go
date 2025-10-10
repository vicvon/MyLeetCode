/*
 * @lc app=leetcode.cn id=77 lang=golang
 * @lcpr version=30203
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	var res [][]int
	var track []int

	var backtrack func(int)
	backtrack = func(start int) {
		if k == len(track) {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
		}

		for i := start; i <= n; i++ {
			track = append(track, i)
			backtrack(i + 1)
			track = track[:len(track)-1]
		}
	}

	backtrack(1)
	return res
}

// @lc code=end

/*
// @lcpr case=start
// 4\n2\n
// @lcpr case=end

// @lcpr case=start
// 1\n1\n
// @lcpr case=end

*/

