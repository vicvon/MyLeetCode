/*
 * @lc app=leetcode.cn id=373 lang=cpp
 *
 * [373] 查找和最小的 K 对数字
 */

// @lc code=start
class Solution
{
public:
    struct node {
        int n1;
        int n2;
        int j;
    };

    vector<vector<int>> kSmallestPairs(vector<int>& nums1, vector<int>& nums2, int k)
    {
        std::vector<std::vector<int>> result;

        std::priority_queue<node, std::vector<node>, std::function<bool(node, node)>> q(
            [](node l, node r) { return (l.n1 + l.n2) > (r.n1 + r.n2); });

        for (int i = 0; i < nums1.size(); i++) {
            q.push({nums1[i], nums2[0], 0});
        }

        while (!q.empty() && k > 0) {
            auto tmp = q.top();
            result.emplace_back(std::vector<int>{tmp.n1, tmp.n2});
            k--;
            q.pop();

            if (tmp.j + 1 < nums2.size()) {
                q.push({tmp.n1, nums2[tmp.j + 1], tmp.j + 1});
            }
        }

        return result;
    }
};
// @lc code=end
