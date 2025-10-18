/*
 * @lc app=leetcode.cn id=303 lang=golang
 * @lcpr version=30203
 *
 * [303] 区域和检索 - 数组不可变
 */

// @lc code=start
type NumArray struct {
	PreSum []int
}

func Constructor(nums []int) NumArray {
	numArray := NumArray{
		PreSum: make([]int, len(nums)+1),
	}
	for i := 1; i <= len(nums); i++ {
		numArray.PreSum[i] = numArray.PreSum[i-1] + nums[i-1]
	}
	return numArray
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.PreSum[right+1] - this.PreSum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// @lc code=end

/*
// @lcpr case=start
// ["NumArray", "sumRange", "sumRange", "sumRange"]\n[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]\n
// @lcpr case=end

*/

