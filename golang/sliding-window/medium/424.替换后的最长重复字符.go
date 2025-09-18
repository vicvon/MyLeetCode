/*
 * @lc app=leetcode.cn id=424 lang=golang
 * @lcpr version=30203
 *
 * [424] 替换后的最长重复字符
 */

// @lc code=start
func characterReplacement(s string, k int) int {
	left, right := 0, 0
	windowCharCount := make([]int, 26)
	windowMaxCount := 0
	res := 0

	for right < len(s) {
		c := s[right] - 'A'
		right++
		windowCharCount[c]++
		if windowCharCount[c] > windowMaxCount {
			windowMaxCount = windowCharCount[c]
		}

		for right-left-windowMaxCount > k {
			windowCharCount[s[left]-'A']--
			left++
		}

		if right-left > res {
			res = right - left
		}
	}

	return res
}

// @lc code=end

/*
// @lcpr case=start
// "ABAB"\n2\n
// @lcpr case=end

// @lcpr case=start
// "AABABBA"\n1\n
// @lcpr case=end

*/

