/*
 * @lc app=leetcode.cn id=1329 lang=golang
 * @lcpr version=30202
 *
 * [1329] 将矩阵按对角线排序
 */

// @lc code=start
func diagonalSort(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])

	diagonals := make(map[int][]int)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			diagonals[i-j] = append(diagonals[i-j], mat[i][j])
		}
	}

	for _, diagonal := range diagonals {
		sort.Slice(diagonal, func(i, j int) bool {
			return diagonal[i] > diagonal[j]
		})
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mat[i][j] = diagonals[i-j][len(diagonals[i-j])-1]
			diagonals[i-j] = diagonals[i-j][:len(diagonals[i-j])-1]
		}
	}

	return mat
}

// @lc code=end

/*
// @lcpr case=start
// [[3,3,1,1],[2,2,1,2],[1,1,1,2]]\n
// @lcpr case=end

// @lcpr case=start
// [[11,25,66,1,69,7],[23,55,17,45,15,52],[75,31,36,44,58,8],[22,27,33,25,68,4],[84,28,14,11,5,50]]\n
// @lcpr case=end

*/

