/*
 * @lc app=leetcode.cn id=977 lang=cpp
 *
 * [977] 有序数组的平方
 */

// @lc code=start
class Solution
{
public:
    vector<int> sortedSquares(vector<int>& nums)
    {
        int              left   = 0;
        int              right  = nums.size() - 1;
        int              cursor = nums.size() - 1;
        std::vector<int> targets(nums.size(), 0);

        while (left <= right) {
            int powLeft  = pow(nums[left], 2);
            int powRight = pow(nums[right], 2);
            if (powLeft < powRight) {
                targets[cursor] = powRight;
                right--;
            } else {
                targets[cursor] = powLeft;
                left++;
            }
            cursor--;
        }
        return targets;
    }
};
// @lc code=end
