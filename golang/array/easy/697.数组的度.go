/*
 * @lc app=leetcode.cn id=697 lang=golang
 * @lcpr version=30203
 *
 * [697] 数组的度
 */

// @lc code=start
func findShortestSubArray(nums []int) int {
	left, right := 0, 0
	target := 0
	targetMap := make(map[int]int)
	numsMap := make(map[int]int)

	for _, v := range nums {
		targetMap[v]++
		if targetMap[v] > target {
			target = targetMap[v]
		}
	}

	// 标记是哪个元素满足了度的要求
	found := make(map[int]int)
	res := len(nums)
	for right < len(nums) {
		num := nums[right]
		numsMap[num]++
		if numsMap[num] == target {
			found[num] = 1
		}
		right++

		for found[num] == 1 && left < right {
			d := nums[left]
			left++
			numsMap[d]--
			// 收缩窗口时,只针对不满足条件的标志重置
			if numsMap[d] != target {
				found[d] = 0
			}

			if res > right-left+1 {
				res = right - left + 1
			}
		}
	}

	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,2,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,2,3,1,4,2]\n
// @lcpr case=end

*/

