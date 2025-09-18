/*
 * @lc app=leetcode.cn id=395 lang=golang
 * @lcpr version=30203
 *
 * [395] 至少有 K 个重复字符的最长子串
 */

// @lc code=start
func longestSubstring(s string, k int) int {
	maxLen := 0
	for i := 1; i <= 26; i++ {
		len := longestKLetterSubstr(s, k, i)
		if len > maxLen {
			maxLen = len
		}
	}

	return maxLen
}

func longestKLetterSubstr(s string, k int, count int) int {
	res := 0
	left, right := 0, 0
	windowCount := make([]int, 26)
	windowUniqueCount := 0
	windowValidCount := 0

	for right < len(s) {
		c := s[right]
		right++
		if windowCount[c-'a'] == 0 {
			windowUniqueCount++
		}
		windowCount[c-'a']++
		if windowCount[c-'a'] == k {
			windowValidCount++
		}

		for windowUniqueCount > count {
			d := s[left]
			left++
			if windowCount[d-'a'] == k {
				windowValidCount--
			}
			windowCount[d-'a']--
			if windowCount[d-'a'] == 0 {
				windowUniqueCount--
			}
		}

		if windowValidCount == count {
			if res < right-left {
				res = right - left
			}
		}
	}

	return res
}

// @lc code=end

/*
// @lcpr case=start
// "aaabb"\n3\n
// @lcpr case=end

// @lcpr case=start
// "ababbc"\n2\n
// @lcpr case=end

*/

