/*
 * @lc app=leetcode.cn id=82 lang=cpp
 *
 * [82] 删除排序链表中的重复元素 II
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
    ListNode *deleteDuplicates(ListNode *head)
    {
        if (head == nullptr || head->next == nullptr) {
            return head;
        }
        auto pre = head;
        auto p   = head->next;

        // 前2个元素不等,直接以第二个元素为head去重
        if (pre->val != p->val) {
            pre->next = deleteDuplicates(p);
            return head;
        }

        // 前2个元素相等,直接删除链表开头重复元素
        auto newHead = deleteFrontDup(pre);
        return deleteDuplicates(newHead);
    }

    ListNode *deleteFrontDup(ListNode *head)
    {
        if (head == nullptr || head->next == nullptr) {
            return head;
        }
        auto pre = head;
        auto p   = head->next;
        if (pre->val != p->val) {
            return pre;
        }
        while (p && (pre->val == p->val)) {
            p = p->next;
        }

        return p;
    }
};
// @lc code=end
