/*
 * @lc app=leetcode.cn id=238 lang=golang
 * @lcpr version=30300
 *
 * [238] 除自身以外数组的乘积
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	preMul := make([]int, len(nums))
	preMul[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		preMul[i] = preMul[i-1] * nums[i]
	}

	sufMul := make([]int, len(nums))
	sufMul[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		sufMul[i] = sufMul[i+1] * nums[i]
	}

	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			ans[i] = sufMul[i+1]
		} else if i == len(nums)-1 {
			ans[i] = preMul[i-1]
		} else {
			ans[i] = preMul[i-1] * sufMul[i+1]
		}
	}

	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [-1,1,0,-3,3]\n
// @lcpr case=end

*/

