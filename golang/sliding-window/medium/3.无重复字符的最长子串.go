/*
 * @lc app=leetcode.cn id=3 lang=golang
 * @lcpr version=30203
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)

	left, right := 0, 0
	res := 0

	for right < len(s) {
		c := s[right]
		right++
		window[c]++

		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}

		if res < right-left {
			res = right - left
		}
	}

	return res
}

// @lc code=end

/*
// @lcpr case=start
// "abcabcbb"\n
// @lcpr case=end

// @lcpr case=start
// "bbbbb"\n
// @lcpr case=end

// @lcpr case=start
// "pwwkew"\n
// @lcpr case=end

*/

