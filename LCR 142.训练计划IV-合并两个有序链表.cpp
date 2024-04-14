/*
 * @lc app=leetcode.cn id=LCR 142 lang=cpp
 *
 * [LCR 142] 训练计划IV
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
    ListNode* trainningPlan(ListNode* l1, ListNode* l2)
    {
        ListNode  dummy;
        ListNode* p  = &dummy;
        auto      p1 = l1;
        auto      p2 = l2;
        while (p1 && p2) {
            if (p1->val < p2->val) {
                p->next = p1;
                p1      = p1->next;
            } else {
                p->next = p2;
                p2      = p2->next;
            }
            p = p->next;
        }

        if (p1 == nullptr) {
            p->next = p2;
        }

        if (p2 == nullptr) {
            p->next = p1;
        }

        return dummy.next;
    }
};