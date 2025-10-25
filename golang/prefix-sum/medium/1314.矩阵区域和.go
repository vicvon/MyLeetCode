/*
 * @lc app=leetcode.cn id=1314 lang=golang
 * @lcpr version=30300
 *
 * [1314] 矩阵区域和
 */

// @lc code=start
func matrixBlockSum(mat [][]int, k int) [][]int {
	numMatrix := NewNumMatrix(mat)

	res := make([][]int, len(mat))
	for i := range res {
		res[i] = make([]int, len(mat[0]))
	}
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			x1 := max(i-k, 0)
			y1 := max(j-k, 0)
			x2 := min(i+k, len(mat)-1)
			y2 := min(j+k, len(mat[0])-1)
			res[i][j] = numMatrix.SumRegion(x1, y1, x2, y2)
		}
	}

	return res
}

type NumMatrix struct {
	PreSum [][]int
}

func NewNumMatrix(matrix [][]int) NumMatrix {
	preSum := make([][]int, len(matrix)+1)
	for i := range preSum {
		preSum[i] = make([]int, len(matrix[0])+1)
	}

	for i := 1; i <= len(matrix); i++ {
		for j := 1; j <= len(matrix[0]); j++ {
			preSum[i][j] = preSum[i][j-1] + preSum[i-1][j] - preSum[i-1][j-1] + matrix[i-1][j-1]
		}
	}

	return NumMatrix{PreSum: preSum}
}

func (m *NumMatrix) SumRegion(x1, y1, x2, y2 int) int {
	return m.PreSum[x2+1][y2+1] - m.PreSum[x2+1][y1] - m.PreSum[x1][y2+1] + m.PreSum[x1][y1]
}

func max(x, y int) int {
	if x >= y {
		return x
	}

	return y
}

func min(x, y int) int {
	if x <= y {
		return x
	}

	return y
}

// @lc code=end

/*
// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n1\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n2\n
// @lcpr case=end

*/

