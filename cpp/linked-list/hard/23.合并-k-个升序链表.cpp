/*
 * @lc app=leetcode.cn id=23 lang=cpp
 *
 * [23] 合并 K 个升序链表
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

        ListNode  dummy;
        ListNode* head = &dummy;

        std::priority_queue<ListNode*, std::vector<ListNode*>, std::function<bool(ListNode*, ListNode*)>> minHeap(
            [](ListNode* l, ListNode* r) { return l->val >= r->val; });

        for (auto list : lists) {
            if (!list)
                continue;
            minHeap.push(list);
        }

        while (!minHeap.empty()) {
            auto node = minHeap.top();
            minHeap.pop();
            head->next = node;
            if (node->next) {
                minHeap.push(node->next);
            }

            head = head->next;
        }

        return dummy.next;
    }
};
// @lc code=end
