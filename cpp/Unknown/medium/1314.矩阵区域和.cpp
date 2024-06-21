/*
 * @lc app=leetcode.cn id=1314 lang=cpp
 *
 * [1314] 矩阵区域和
 */

// @lc code=start
class Solution
{
public:
    class NumMatrix
    {
    private:
        vector<vector<int>> preSum;

    public:
        NumMatrix(vector<vector<int>>& matrix)
        {
            int m = matrix.size();
            int n = matrix[0].size();
            if (m == 0 || n == 0)
                return;

            preSum = vector<vector<int>>(m + 1, vector<int>(n + 1));
            for (int i = 1; i <= m; i++) {
                for (int j = 1; j <= n; j++) {
                    preSum[i][j] = preSum[i - 1][j] + preSum[i][j - 1] + matrix[i - 1][j - 1] - preSum[i - 1][j - 1];
                }
            }
        }

        int sumRegion(int x1, int y1, int x2, int y2)
        {
            return preSum[x2 + 1][y2 + 1] - preSum[x2 + 1][y1] - preSum[x1][y2 + 1] + preSum[x1][y1];
        }
    };

    vector<vector<int>> matrixBlockSum(vector<vector<int>>& mat, int k)
    {
        NumMatrix           numMatrix(mat);
        int                 m = mat.size();
        int                 n = mat[0].size();
        vector<vector<int>> res(m, vector<int>(n));
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                int x1    = max(i - k, 0);
                int y1    = max(j - k, 0);
                int x2    = min(i + k, m - 1);
                int y2    = min(j + k, n - 1);
                res[i][j] = numMatrix.sumRegion(x1, y1, x2, y2);
            }
        }

        return res;
    }
};
// @lc code=end
