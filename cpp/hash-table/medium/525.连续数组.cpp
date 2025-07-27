/*
 * @lc app=leetcode.cn id=525 lang=cpp
 *
 * [525] 连续数组
 */

// @lc code=start
class Solution
{
public:
    int findMaxLength(vector<int>& nums)
    {
#if 0
        // preSum是数组
        vector<int>             preSum(nums.size() + 1, 0);
        unordered_map<int, int> valToIdx;
        for (int i = 1; i < nums.size() + 1; i++) {
            preSum[i] = preSum[i - 1] + (nums[i - 1] == 0 ? -1 : 1);
        }

        for (int i = 0; i < preSum.size(); i++) {
            valToIdx[preSum[i]] = i;
        }

        int maxLen = 0;
        for (int i = 0; i < preSum.size(); i++) {
            int need = preSum[i];
            if (valToIdx.find(need) != valToIdx.end()) {
                int len = abs(valToIdx[need] - i);
                if (maxLen < len) {
                    maxLen = len;
                }
            }
        }
#else
        // preSum是变量
        unordered_map<int, int> valToIdx;
        int                     preSum = 0;
        int                     maxLen = 0;
        for (int i = 0; i < nums.size(); i++) {
            if (valToIdx.find(preSum) == valToIdx.end()) {
                valToIdx[preSum] = 0;
            } else {
                maxLen = max(maxLen, abs(valToIdx[preSum] - i) + 1);
                preSum = preSum + (nums[i] == 0 ? -1 : 1);
            }
        }
#endif

        return maxLen;
    }
};
// @lc code=end
