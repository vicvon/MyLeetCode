/*
 * @lc app=leetcode.cn id=867 lang=golang
 * @lcpr version=30202
 *
 * [867] 转置矩阵
 */

// @lc code=start
func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[j][i] = matrix[i][j]
		}
	}

	return res
}

// @lc code=end

/*
// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3],[4,5,6]]\n
// @lcpr case=end

*/

