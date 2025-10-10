/*
 * @lc app=leetcode.cn id=90 lang=golang
 * @lcpr version=30203
 *
 * [90] 子集 II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	var track []int

	sort.Ints(nums)
	var backtrack func(int)
	backtrack = func(start int) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		res = append(res, tmp)

		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			track = append(track, nums[i])
			backtrack(i + 1)
			track = track[:len(track)-1]
		}
	}

	backtrack(0)
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,2]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

*/

