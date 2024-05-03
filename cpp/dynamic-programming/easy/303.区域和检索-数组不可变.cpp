/*
 * @lc app=leetcode.cn id=303 lang=cpp
 *
 * [303] 区域和检索 - 数组不可变
 */

// @lc code=start
class NumArray
{
public:
    vector<int> vec_;
    vector<int> preSum_;
    NumArray(vector<int>& nums)
    {
        vec_ = nums;
        preSum_.resize(nums.size() + 1);
        for (int i = 1; i < preSum_.size(); i++) {
            preSum_[i] = preSum_[i - 1] + nums[i - 1];
        }
    }

    int sumRange(int left, int right)
    {
        return preSum_[right + 1] - preSum_[left];
    }
};

/**
 * Your NumArray object will be instantiated and called as such:
 * NumArray* obj = new NumArray(nums);
 * int param_1 = obj->sumRange(left,right);
 */
// @lc code=end
