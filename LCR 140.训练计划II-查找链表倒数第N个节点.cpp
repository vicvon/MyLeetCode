/*
 * @lc app=leetcode.cn id=LCR 140 lang=cpp
 *
 * [LCR 140] 训练计划II
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
    ListNode* trainingPlan(ListNode* head, int cnt)
    {
        auto p1 = head;
        auto p2 = head;

        for (int i = 0; i < cnt; i++) {
            p2 = p2->next;
        }

        while (p2) {
            p1 = p1->next;
            p2 = p2->next;
        }

        return p1;
    }
};