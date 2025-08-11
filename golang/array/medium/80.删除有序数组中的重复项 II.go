/*
 * @lc app=leetcode.cn id=80 lang=golang
 * @lcpr version=30202
 *
 * [80] 删除有序数组中的重复项 II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	slow, fast := 0, 0
	count := 0
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		} else if slow < fast && count < 2 {
			slow++
			nums[slow] = nums[fast]
		}

		fast++
		count++
		if fast < len(nums) && nums[fast] != nums[fast-1] {
			count = 0
		}
	}

	return slow + 1
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,1,2,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [0,0,1,1,1,1,2,3,3]\n
// @lcpr case=end

*/

