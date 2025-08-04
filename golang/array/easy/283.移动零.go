/*
 * @lc app=leetcode.cn id=283 lang=golang
 * @lcpr version=30202
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}

		fast++
	}

	for slow < len(nums) {
		nums[slow] = 0
		slow++
	}
}

// @lc code=end

/*
// @lcpr case=start
// [0,1,0,3,12]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

*/

