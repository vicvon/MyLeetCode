/*
 * @lc app=leetcode.cn id=59 lang=golang
 * @lcpr version=30202
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	upperBound := 0
	lowerBound := n - 1
	leftBound := 0
	rightBound := n - 1
	num := 1

	for num <= n*n {
		if upperBound <= lowerBound {
			for i := leftBound; i <= rightBound; i++ {
				matrix[upperBound][i] = num
				num++
			}
			upperBound++
		}

		if rightBound >= leftBound {
			for i := upperBound; i <= lowerBound; i++ {
				matrix[i][rightBound] = num
				num++
			}
			rightBound--
		}

		if lowerBound >= upperBound {
			for i := rightBound; i >= leftBound; i-- {
				matrix[lowerBound][i] = num
				num++
			}
			lowerBound--
		}

		if leftBound <= rightBound {
			for i := lowerBound; i >= upperBound; i-- {
				matrix[i][leftBound] = num
				num++
			}
			leftBound++
		}
	}

	return matrix
}

// @lc code=end

/*
// @lcpr case=start
// 3\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/

