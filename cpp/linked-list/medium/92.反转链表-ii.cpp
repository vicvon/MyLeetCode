/*
 * @lc app=leetcode.cn id=92 lang=cpp
 *
 * [92] 反转链表 II
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
    ListNode *reverseBetween(ListNode *head, int left, int right)
    {
        ListNode dummy;
        dummy.next = head;
        auto p     = head;
        auto pre   = &dummy;
        for (int i = 1; i < left; i++) {
            pre = p;
            p   = p->next;
        }
        auto h    = reverseNList(p, right - left + 1);
        pre->next = h;

        return dummy.next;
    }

    ListNode *reverseNList(ListNode *head, int n)
    {
        if (n == 1) {
            node = head->next;
            return head;
        }

        auto newHead     = reverseNList(head->next, n - 1);
        head->next->next = head;
        head->next       = node;
        return newHead;
    }
};
// @lc code=end
