/*
 * @lc app=leetcode.cn id=283 lang=cpp
 *
 * [283] 移动零
 */

// @lc code=start
class Solution
{
public:
    void moveZeroes(vector<int>& nums)
    {
        if (nums.size() <= 1)
            return;
        int slow = 0;
        int fast = 0;
        while (fast < nums.size()) {
            if (nums[fast] != 0) {
                nums[slow] = nums[fast];
                // nums[fast] = 0;
                slow++;
            }
            fast++;
        }
        while (slow < nums.size()) {
            nums[slow] = 0;
            slow++;
        }
    }
};
// @lc code=end
