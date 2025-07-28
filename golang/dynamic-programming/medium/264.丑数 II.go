/*
 * @lc app=leetcode.cn id=264 lang=golang
 * @lcpr version=30202
 *
 * [264] 丑数 II
 */

// @lc code=start
func nthUglyNumber(n int) int {
	ugly := make([]int, n)
	ugly[0] = 1
	i2, i3, i5 := 0, 0, 0
	p := 1
	for p < n {
		next2, next3, next5 := ugly[i2]*2, ugly[i3]*3, ugly[i5]*5
		next := min(next2, min(next3, next5))
		ugly[p] = next
		p++
		if next == next2 {
			i2++
		}
		if next == next3 {
			i3++
		}
		if next == next5 {
			i5++
		}
	}

	return ugly[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// 10\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/

