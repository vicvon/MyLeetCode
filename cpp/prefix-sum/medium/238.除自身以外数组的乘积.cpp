/*
 * @lc app=leetcode.cn id=238 lang=cpp
 *
 * [238] 除自身以外数组的乘积
 */

// @lc code=start
class Solution
{
public:
    vector<int> productExceptSelf(vector<int>& nums)
    {
        vector<int> preMul(nums.size(), nums[0]);
        for (int i = 1; i < nums.size(); i++) {
            preMul[i] = preMul[i - 1] * nums[i];
        }

        vector<int> sufMul(nums.size(), nums[nums.size() - 1]);
        for (int i = nums.size() - 2; i >= 0; i--) {
            sufMul[i] = sufMul[i + 1] * nums[i];
        }

        vector<int> ans(nums.size(), 0);
        for (int i = 0; i < nums.size(); i++) {
            if (i == 0) {
                ans[i] = sufMul[1];
                continue;
            } else if (i == nums.size() - 1) {
                ans[i] = preMul[i - 1];
                continue;
            }

            ans[i] = preMul[i - 1] * sufMul[i + 1];
        }

        return ans;
    }
};
// @lc code=end
