/*
 * @lc app=leetcode.cn id=373 lang=golang
 * @lcpr version=30202
 *
 * [373] 查找和最小的 K 对数字
 */

// @lc code=start
type MinHeap [][]int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return (h[i][0] + h[i][1]) < (h[j][0] + h[j][1])
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

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var h MinHeap
	heap.Init(&h)

	for _, row := range nums1 {
		heap.Push(&h, []int{row, nums2[0], 0})
	}

	res := [][]int{}
	for h.Len() > 0 && k > 0 {
		top := heap.Pop(&h).([]int)
		k--
		nextPos := top[2] + 1
		if nextPos < len(nums2) {
			heap.Push(&h, []int{top[0], nums2[nextPos], nextPos})
		}
		res = append(res, []int{top[0], top[1]})
	}

	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,7,11]\n[2,4,6]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,1,2]\n[1,2,3]\n2\n
// @lcpr case=end

*/

