/*
 * @lc app=leetcode.cn id=378 lang=golang
 * @lcpr version=30202
 *
 * [378] 有序矩阵中第 K 小的元素
 */

// @lc code=start
type MinHeap [][]int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kthSmallest(matrix [][]int, k int) int {
	var h MinHeap
	heap.Init(&h)
	for i, row := range matrix {
		if len(row) > 0 {
			h = append(h, []int{row[0], i, 0}) // value, row index, column index
		}
	}

	res := 0
	for k > 0 {
		top := heap.Pop(&h).([]int)
		if k == 1 {
			res = top[0]
			break
		}
		k--
		row, col := top[1], top[2]
		if col+1 < len(matrix[row]) {
			heap.Push(&h, []int{matrix[row][col+1], row, col + 1})
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [[1,5,9],[10,11,13],[12,13,15]]\n8\n
// @lcpr case=end

// @lcpr case=start
// [[-5]]\n1\n
// @lcpr case=end

*/

