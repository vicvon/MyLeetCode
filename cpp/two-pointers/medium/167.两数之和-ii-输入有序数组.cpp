/*
 * @lc app=leetcode.cn id=167 lang=cpp
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start
class Solution
{
public:
    vector<int> twoSum(vector<int>& numbers, int target)
    {
        std::vector<int> result;
        int              left  = 0;
        int              right = numbers.size() - 1;
        while (left <= right) {
            int number = numbers[left] + numbers[right];
            if (number == target) {
                result.push_back(left + 1);
                result.push_back(right + 1);
                break;
            }
            if (number < target) {
                left++;
            } else {
                right--;
            }
        }
        return result;
    }
};
// @lc code=end
