/*
 * @lc app=leetcode.cn id=75 lang=golang
 * @lcpr version=30202
 *
 * [75] 颜色分类
 */

// @lc code=start
func sortColors(nums []int) {
	p0, p2 := 0, len(nums)-1
	p1 := 0
	for p1 <= p2 {
		if nums[p1] == 0 {
			nums[p0], nums[p1] = nums[p1], nums[p0]
			p0++
		} else if nums[p1] == 2 {
			nums[p2], nums[p1] = nums[p1], nums[p2]
			p2--
		} else if nums[p1] == 1 {
			p1++
		}

		if p1 < p0 {
			p1 = p0
		}
	}
}

// @lc code=end

/*
// @lcpr case=start
// [2,0,2,1,1,0]\n
// @lcpr case=end

// @lcpr case=start
// [2,0,1]\n
// @lcpr case=end

*/

