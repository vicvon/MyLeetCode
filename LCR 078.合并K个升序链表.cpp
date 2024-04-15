/*
 * @lc app=leetcode.cn id=LCR 142 lang=cpp
 *
 * [LCR 078] 合并K个升序链表
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
    ListNode* mergeKLists(vector<ListNode*>& lists)
    {
        if (lists.empty()) {
            return nullptr;
        }

        std::priority_queue<ListNode*, std::vector<ListNode*>, std::function<bool(ListNode*, ListNode*)>> pq(
            [](ListNode* l, ListNode* r) { return l->val > r->val; });

        for (auto l : lists) {
            if (l == nullptr) {
                continue;
            }
            pq.push(l);
        }

        ListNode dummy;
        auto     p = &dummy;
        while (!pq.empty()) {
            auto node = pq.top();
            p->next   = node;
            pq.pop();

            if (node->next) {
                pq.push(node->next);
            }
            p = p->next;
        }

        return dummy.next;
    }
};