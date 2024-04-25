/*
[https://labuladong.online/algo/problem-set/array-two-pointers/#%E9%A2%98%E7%9B%AE%E6%8F%8F%E8%BF%B0-2]
给你一个已经 排好序 的整数数组 nums 和整数 a 、 b 、 c 。对于数组中的每一个元素 nums[i] ，计算函数值 f(x) = ax2 + bx + c
，请 按升序返回数组 。

示例 1：

输入: nums = [-4,-2,2,4], a = 1, b = 3, c = 5
输出: [3,9,15,33]
示例 2：

输入: nums = [-4,-2,2,4], a = -1, b = 3, c = 5
输出: [-23,-5,1,7]
提示：

1 <= nums.length <= 200
-100 <= nums[i], a, b, c <= 100
nums 按照 升序排列
进阶：你可以在时间复杂度为 O(n) 的情况下解决这个问题吗？
*/
class Solution
{
public:
    vector<int> sortTransformedArray(vector<int>& nums, int a, int b, int c)
    {
        int              left   = 0;
        int              right  = nums.size() - 1;
        int              cursor = a > 0 ? nums.size() - 1 : 0;
        std::vector<int> targets(nums.size(), 0);

        while (left <= right) {
            int fLeft  = f(nums[left], a, b, c);
            int fRigth = f(nums[right], a, b, c);
            if (a > 0) {
                if (fLeft < fRigth) {
                    targets[cursor] = fRigth;
                    right--;
                } else {
                    targets[cursor] = fLeft;
                    left++;
                }
                cursor--;
            } else {
                if (fLeft < fRigth) {
                    targets[cursor] = fLeft;
                    right--;
                } else {
                    targets[cursor] = fRigth;
                    left++;
                }
                cursor++;
            }
        }
        return targets;
    }

    int f(int x, int a, int b, int c)
    {
        return a * x * x + b * x + c;
    }
};