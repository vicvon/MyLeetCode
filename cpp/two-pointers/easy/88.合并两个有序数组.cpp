/*
 * @lc app=leetcode.cn id=88 lang=cpp
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
class Solution
{
public:
    void merge(vector<int>& nums1, int m, vector<int>& nums2, int n)
    {
        int n1 = m - 1;
        int n2 = n - 1;
        int p  = nums1.size() - 1;
        while (n1 >= 0 && n2 >= 0) {
            if (nums1[n1] >= nums2[n2]) {
                nums1[p] = nums1[n1];
                n1--;
            } else {
                nums1[p] = nums2[n2];
                n2--;
            }
            p--;
        }

        while (n2 >= 0) {
            nums1[p] = nums2[n2];
            p--;
            n2--;
        }
    }
};
// @lc code=end
