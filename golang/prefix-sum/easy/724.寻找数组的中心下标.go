/*
 * @lc app=leetcode.cn id=724 lang=golang
 * @lcpr version=30300
 *
 * [724] 寻找数组的中心下标
 */

// @lc code=start
func pivotIndex(nums []int) int {
	numArr := NewNumArray(nums)
	for i := 0; i < len(nums); i++ {
		left, right := -1, 1
		if i == 0 {
			left = 0
			right = numArr.SumRange(i+1, len(nums)-1)
		} else if i == len(nums)-1 {
			left = numArr.SumRange(0, i-1)
			right = 0
		} else {
			left = numArr.SumRange(0, i-1)
			right = numArr.SumRange(i+1, len(nums)-1)
		}

		if left == right {
			return i
		}
	}

	return -1
}

type NumArray struct {
	PreSum []int
}

func NewNumArray(nums []int) NumArray {
	preSum := make([]int, len(nums)+1)

	for i := 1; i <= len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	return NumArray{PreSum: preSum}
}

func (n *NumArray) SumRange(i, j int) int {
	return n.PreSum[j+1] - n.PreSum[i]
}

// @lc code=end

/*
// @lcpr case=start
// [1, 7, 3, 6, 5, 6]\n
// @lcpr case=end

// @lcpr case=start
// [1, 2, 3]\n
// @lcpr case=end

// @lcpr case=start
// [2, 1, -1]\n
// @lcpr case=end

*/

