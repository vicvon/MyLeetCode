/*
 * @lc app=leetcode.cn id=378 lang=cpp
 *
 * [378] 有序矩阵中第 K 小的元素
 */

// @lc code=start
class Solution
{
public:
    struct node {
        int val;
        int index;
    };

    int kthSmallest(vector<vector<int>>& matrix, int k)
    {
        int target = 0;

        std::vector<int> pos(matrix.size(), 1);

        std::priority_queue<node, std::vector<node>, std::function<bool(node, node)>> q(
            [](node l, node r) { return l.val > r.val; });
        for (int i = 0; i < matrix.size(); i++) {
            if (matrix[i].empty())
                continue;
            q.push({matrix[i][0], i});
        }

        while (!q.empty()) {
            auto tmp = q.top();
            if (k == 1) {
                target = tmp.val;
                break;
            }

            k--;
            q.pop();
            if (pos[tmp.index] < matrix[tmp.index].size()) {
                q.push({matrix[tmp.index][pos[tmp.index]++], tmp.index});
            }
        }

        return target;
    }
};
// @lc code=end
