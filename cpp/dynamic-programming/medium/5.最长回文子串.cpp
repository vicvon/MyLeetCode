/*
 * @lc app=leetcode.cn id=5 lang=cpp
 *
 * [5] 最长回文子串
 */

// @lc code=start
class Solution
{
public:
    string longestPalindrome(string s)
    {
        string res = "";
        for (int i = 0; i < s.length(); i++) {
            auto s1 = palindrome(s, i, i);
            auto s2 = palindrome(s, i, i + 1);

            res = res.length() > s1.length() ? res : s1;
            res = res.length() > s2.length() ? res : s2;
        }

        return res;
    }

    string palindrome(string s, int left, int right)
    {
        while (left >= 0 && right <= s.length() - 1 && s[left] == s[right]) {
            left--;
            right++;
        }

        return s.substr(left + 1, right - left - 1);
    }
};
// @lc code=end
