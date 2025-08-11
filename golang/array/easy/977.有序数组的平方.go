/*
 * @lc app=leetcode.cn id=977 lang=golang
 * @lcpr version=30202
 *
 * [977] 有序数组的平方
 */

// @lc code=start
func sortedSquares(nums []int) []int {
	left, right := 0, len(nums)-1
	res := make([]int, len(nums))
	pos := len(nums) - 1
	for left <= right {
		if nums[left]*nums[left] > nums[right]*nums[right] {
			res[pos] = nums[left] * nums[left]
			left++
		} else {
			res[pos] = nums[right] * nums[right]
			right--
		}
		pos--
	}

	return res
}

// @lc code=end

/*
// @lcpr case=start
// [-4,-1,0,3,10]\n
// @lcpr case=end

// @lcpr case=start
// [-7,-3,2,3,11]\n
// @lcpr case=end

*/

