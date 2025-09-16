/*
 * @lc app=leetcode.cn id=567 lang=golang
 * @lcpr version=30203
 *
 * [567] 字符串的排列
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	need, window := make(map[byte]int), make(map[byte]int)
	for i := range s1 {
		need[s1[i]]++
	}

	left, right := 0, 0
	valid := 0
	for right < len(s2) {
		c := s2[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}
			d := s2[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	return false
}

// @lc code=end

/*
// @lcpr case=start
// "eidbaooo"\n
// @lcpr case=end

// @lcpr case=start
// "eidboaoo"\n
// @lcpr case=end

*/

