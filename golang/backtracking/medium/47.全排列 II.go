/*
 * @lc app=leetcode.cn id=47 lang=golang
 * @lcpr version=30203
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	var res [][]int
	var track []int
	visited := make([]bool, len(nums))

	sort.Ints(nums)
	var backtrack func()
	backtrack = func() {
		if len(track) == len(nums) {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)

			return
		}

		for i := 0; i < len(nums); i++ {
			if visited[i] {
				continue
			}

			if i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
				continue
			}

			visited[i] = true
			track = append(track, nums[i])

			backtrack()

			track = track[:len(track)-1]
			visited[i] = false
		}
	}

	backtrack()
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

*/

