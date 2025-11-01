/*
 * @lc app=leetcode.cn id=1352 lang=golang
 * @lcpr version=30300
 *
 * [1352] 最后 K 个数的乘积
 */

// @lc code=start
type ProductOfNumbers struct {
	preProduct []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{preProduct: []int{1}}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.preProduct = []int{1}
		return
	}

	n := len(this.preProduct)
	this.preProduct = append(this.preProduct, this.preProduct[n-1]*num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	n := len(this.preProduct)
	if n-1 < k {
		return 0
	}

	return this.preProduct[n-1] / this.preProduct[n-k-1]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
// @lc code=end

/*
// @lcpr case=start
// ["ProductOfNumbers","add","add","add","add","add","getProduct","getProduct","getProduct","add","getProduct"]\n[[],[3],[0],[2],[5],[4],[2],[3],[4],[8],[2]]\n
// @lcpr case=end

*/

