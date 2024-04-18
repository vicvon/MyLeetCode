/*
 * @lc app=leetcode.cn id=25 lang=cpp
 *
 * [25] K 个一组翻转链表
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
    ListNode *node = nullptr;
    ListNode *reverseKGroup(ListNode *head, int k)
    {
        auto p = head;
        for (int i = 0; i < k; i++)
        {
            // 不足k个元素,不翻转
            if (p == nullptr)
            {
                return head;
            }
            p = p->next;
        }
        // 翻转head之后的k个元素
        auto newHead = reverseNList(head, k);
        // 翻转之后head变尾结点,链接剩下的链表k组翻转结果
        head->next = reverseKGroup(p, k);
        return newHead;
    }

    ListNode *reverseNList(ListNode *head, int n)
    {
        if (n == 1)
        {
            node = head->next;
            return head;
        }

        auto newHead = reverseNList(head->next, n - 1);
        head->next->next = head;
        head->next = node;
        return newHead;
    }
};
// @lc code=end
