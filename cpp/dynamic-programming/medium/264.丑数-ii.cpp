/*
 * @lc app=leetcode.cn id=264 lang=cpp
 *
 * [264] 丑数 II
 */

// @lc code=start
class Solution
{
public:
    int nthUglyNumber(int n)
    {
        std::vector<int> ugly(n, 0);

        int p2 = 0, p3 = 0, p5 = 0;

        int p   = 1;
        ugly[0] = 1;
        while (p < n) {
            ugly[p] = min(min(ugly[p2] * 2, ugly[p3] * 3), ugly[p5] * 5);
            if (ugly[p] == ugly[p2] * 2)
                p2++;
            if (ugly[p] == ugly[p3] * 3)
                p3++;
            if (ugly[p] == ugly[p5] * 5)
                p5++;
            p++;
        }

        return ugly[p - 1];
    }
};
// @lc code=end
