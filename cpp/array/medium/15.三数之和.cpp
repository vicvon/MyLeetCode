/*
 * @lc app=leetcode.cn id=15 lang=cpp
 *
 * [15] 三数之和
 */

// @lc code=start
class Solution
{
public:
    vector<vector<int>> threeSum(vector<int>& nums)
    {
        vector<vector<int>> res;
        sort(nums.begin(), nums.end());
        for (int i = 0; i < nums.size() - 2;) {
            int  target = 0 - nums[i];
            auto result = twoSum(nums, i + 1, target);
            for (auto e : result) {
                e.push_back(nums[i]);
                res.push_back(e);
            }

            int t = nums[i];
            while ((i < nums.size() - 2) && t == nums[i])
                i++;
        }

        return res;
    }

    vector<vector<int>> twoSum(vector<int>& nums, int start, int target)
    {
        vector<vector<int>> result;
        int                 left = start, right = nums.size() - 1;
        while (left < right) {
            int sum  = nums[left] + nums[right];
            int lNum = nums[left];
            int rNum = nums[right];
            if (sum < target) {
                while (left < right && lNum == nums[left]) {
                    left++;
                }
            } else if (sum > target) {
                while (left < right && rNum == nums[right]) {
                    right--;
                }
            } else {
                result.push_back({nums[left], nums[right]});
                while (left < right && lNum == nums[left]) {
                    left++;
                }
                while (left < right && rNum == nums[right]) {
                    right--;
                }
            }
        }
        return result;
    }
};
// @lc code=end
