/*
 * @lc app=leetcode.cn id=54 lang=golang
 * @lcpr version=30202
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	var res []int
	rowNum := len(matrix)
	columNum := len(matrix[0])
	upperBound := 0
	leftBound := 0
	rightBound := columNum - 1
	lowerBound := rowNum - 1

	for len(res) < rowNum*columNum {
		if upperBound <= lowerBound {
			for j := leftBound; j <= rightBound; j++ {
				res = append(res, matrix[upperBound][j])
			}
			upperBound++
		}

		if rightBound >= leftBound {
			for j := upperBound; j <= lowerBound; j++ {
				res = append(res, matrix[j][rightBound])
			}
			rightBound--
		}

		if lowerBound >= upperBound {
			for j := rightBound; j >= leftBound; j-- {
				res = append(res, matrix[lowerBound][j])
			}
			lowerBound--
		}

		if leftBound <= rightBound {
			for j := lowerBound; j >= upperBound; j-- {
				res = append(res, matrix[j][leftBound])
			}
			leftBound++
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
// [[1,2,3,4],[5,6,7,8],[9,10,11,12]]\n
// @lcpr case=end

*/

