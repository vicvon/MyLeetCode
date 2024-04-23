/*
 * @lc app=leetcode.cn id=26 lang=cpp
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
class Solution
{
public:
    int removeDuplicates(vector<int> &nums)
    {
        if (nums.empty())
            return 0;

        int start = 0;
        for (int i = 1; i < nums.size(); i++) {
            if (nums[start] == nums[i]) {
                continue;
            } else {
                nums[start + 1] = nums[i];
                start++;
            }
        }

        return start + 1;
    }
};
// @lc code=end
