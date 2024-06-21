/*
 * @lc app=leetcode.cn id=724 lang=cpp
 *
 * [724] 寻找数组的中心下标
 */

// @lc code=start
class Solution
{
public:
    int pivotIndex(vector<int>& nums)
    {
        vector<int> preSum(nums.size() + 1);
        for (int i = 1; i <= nums.size(); i++) {
            preSum[i] = preSum[i - 1] + nums[i - 1];
        }
        int pos  = -1;
        int left = 0, right = 0;
        for (int i = 0; i < nums.size(); i++) {
            left  = preSum[i];
            right = preSum[nums.size()] - preSum[i + 1];

            if (left == right) {
                pos = i;
                break;
            }
        }
        return pos;
    }
};
// @lc code=end
