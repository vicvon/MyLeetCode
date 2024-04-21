/*
 * @lc app=leetcode.cn id=234 lang=cpp
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution
{
public:
    ListNode *left = nullptr;

    bool isPalindrome(ListNode *head)
    {
        left = head;
        return traverse(head);
    }

    bool traverse(ListNode *right)
    {
        if (right == nullptr) {
            return true;
        }
        // 后序递归遍历
        auto res = traverse(right->next);
        res      = res && (left->val == right->val);
        left     = left->next;
        return res;
    }
};
// @lc code=end
