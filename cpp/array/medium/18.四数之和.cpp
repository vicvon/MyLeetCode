/*
 * @lc app=leetcode.cn id=18 lang=cpp
 *
 * [18] 四数之和
 */

// @lc code=start
class Solution
{
public:
    vector<vector<int>> fourSum(vector<int>& nums, int target)
    {
        vector<vector<int>> res;
        sort(nums.begin(), nums.end());
        for (int i = 0; i < nums.size(); i++) {
            auto result = threeSum(nums, i + 1, target - nums[i]);
            for (auto e : result) {
                e.push_back(nums[i]);
                res.push_back(e);
            }

            while ((i < nums.size() - 1) && nums[i] == nums[i + 1])
                i++;
        }

        return res;
    }

    vector<vector<int>> threeSum(vector<int>& nums, int start, long target)
    {
        vector<vector<int>> res;
        for (int i = start; i < nums.size();) {
            auto result = twoSum(nums, i + 1, target - nums[i]);
            for (auto e : result) {
                e.push_back(nums[i]);
                res.push_back(e);
            }

            int t = nums[i];
            while ((i < nums.size()) && t == nums[i])
                i++;
        }

        return res;
    }

    vector<vector<int>> twoSum(vector<int>& nums, int start, long target)
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
