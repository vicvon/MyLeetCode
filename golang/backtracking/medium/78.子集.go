/*
 * @lc app=leetcode.cn id=78 lang=golang
 * @lcpr version=30203
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	var res [][]int
	var track []int

	var backtrace func(int)
	backtrace = func(start int) {
		res = append(res, append([]int(nil), track...))

		for i := start; i < len(nums); i++ {
			track = append(track, nums[i])
			backtrace(i + 1)
			track = track[:len(track)-1]
		}
	}

	backtrace(0)
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

*/

