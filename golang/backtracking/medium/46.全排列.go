/*
 * @lc app=leetcode.cn id=46 lang=golang
 * @lcpr version=30203
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	var res [][]int
	var track []int
	visited := make([]int, len(nums))

	var backtrack func()
	backtrack = func() {
		if len(track) == len(nums) {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if visited[i] != 0 {
				continue
			}

			track = append(track, nums[i])
			visited[i] = 1
			backtrack()
			track = track[:len(track)-1]
			visited[i] = 0
		}
	}

	backtrack()
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [0,1]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/

