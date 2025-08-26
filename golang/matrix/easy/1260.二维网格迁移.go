/*
 * @lc app=leetcode.cn id=1260 lang=golang
 * @lcpr version=30202
 *
 * [1260] 二维网格迁移
 */

// @lc code=start
func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	mn := m * n
	k = k % mn
	arr := make([]int, mn)
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if count == mn-k {
				count = -k
			}
			arr[k+count] = grid[i][j]
			count++
		}
	}

	pos := 0
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			res[i][j] = arr[pos]
			pos++
		}
	}

	return res
}

// @lc code=end

/*
// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n1\n
// @lcpr case=end

// @lcpr case=start
// [[3,8,1,9],[19,7,2,5],[4,6,11,10],[12,0,21,13]]\n4\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n9\n
// @lcpr case=end

*/

