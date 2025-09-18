/*
 * @lc app=leetcode.cn id=220 lang=golang
 * @lcpr version=30203
 *
 * [220] 存在重复元素 III
 */

// @lc code=start
func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	if indexDiff <= 0 || valueDiff < 0 {
		return false
	}

	getID := func(x, w int) int {
		if x >= 0 {
			return x / w
		}

		return (x+1)/w - 1
	}

	window := make(map[int]int)
	w := valueDiff + 1

	for i := 0; i < len(nums); i++ {
		m := getID(nums[i], w)

		if _, ok := window[m]; ok {
			return true
		}

		if v, ok := window[m-1]; ok && math.Abs(float64(nums[i]-v)) < float64(w) {
			return true
		}
		if v, ok := window[m+1]; ok && math.Abs(float64(nums[i]-v)) < float64(w) {
			return true
		}

		window[m] = nums[i]
		if i >= indexDiff {
			delete(window, getID(nums[i-indexDiff], w))
		}
	}

	return false
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,1]\n3\n0\n
// @lcpr case=end

// @lcpr case=start
// [1,5,9,1,5,9]\n2\n3\n
// @lcpr case=end

*/

