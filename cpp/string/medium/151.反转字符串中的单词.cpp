/*
 * @lc app=leetcode.cn id=151 lang=cpp
 *
 * [151] 反转字符串中的单词
 */

// @lc code=start
class Solution
{
public:
    string reverseWords(string s)
    {
        string res;
        for (int i = 0; i < s.length(); i++) {
            char c = s[i];
            if (c != ' ') {
                res += c;
            } else if (res.length() > 0 && res.back() != ' ') {
                res += ' ';
            }
        }

        if (res.back() == ' ') {
            res.pop_back();
        }

        int start = -1, end = 0;
        for (int i = 0; i < res.length(); i++) {
            if (res[i] == ' ') {
                res   = reverseString(res, start + 1, i - 1);
                start = i;
            }
        }

        if (start == -1) {
            return res;
        }

        if (start < res.length() - 1) {
            res = reverseString(res, start + 1, res.length() - 1);
        }

        return reverseString(res, 0, res.length() - 1);
    }

    string reverseString(string s, int left, int right)
    {
        while (left <= right) {
            char tmp = s[left];
            s[left]  = s[right];
            s[right] = tmp;
            left++;
            right--;
        }
        return s;
    }
};
// @lc code=end
