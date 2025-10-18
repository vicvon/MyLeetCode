/*
 * @lc app=leetcode.cn id=1109 lang=golang
 * @lcpr version=30203
 *
 * [1109] 航班预订统计
 */

// @lc code=start
func corpFlightBookings(bookings [][]int, n int) []int {
	result := make([]int, n)
	diff := make([]int, n)

	for _, booking := range bookings {
		left := booking[0] - 1
		right := booking[1] - 1
		val := booking[2]

		diff[left] += val
		if right+1 < len(diff) {
			diff[right+1] -= val
		}
	}

	result[0] = diff[0]
	for i := 1; i < len(diff); i++ {
		result[i] = result[i-1] + diff[i]
	}

	return result
}

// @lc code=end

/*
// @lcpr case=start
// [[1,2,10],[2,3,20],[2,5,25]]\n5\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,10],[2,2,15]]\n2\n
// @lcpr case=end

*/

