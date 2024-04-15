/*
 * @lc app=leetcode.cn id=1 lang=cpp
 *
 * [1] 两数之和
 */

// @lc code=start
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        std::unordered_map<int, int> hashMap;
        for (int i = 0; i < nums.size(); i++ ) {
            int need = target - nums[i];
            if (hashMap.find(need) != hashMap.end()) {
                return vector<int>{i, hashMap[need]};
            }

            hashMap.insert(pair<int, int>(nums[i], i));
        }

        return vector<int>();
    }
};
// @lc code=end

