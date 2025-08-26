/*
 * @lc app=leetcode.cn id=14 lang=golang
 * @lcpr version=30202
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	m := len(strs)
	n := len(strs[0])

	for col := 0; col < n; col++ {
		for row := 1; row < m; row++ {
			thisStr, preStr := strs[row], strs[row-1]
			if col >= len(thisStr) || col >= len(preStr) || thisStr[col] != preStr[col] {
				return strs[row][:col]
			}
		}
	}

	return strs[0]
}

// @lc code=end

/*
// @lcpr case=start
// ["flower","flow","flight"]\n
// @lcpr case=end

// @lcpr case=start
// ["dog","racecar","car"]\n
// @lcpr case=end

*/

