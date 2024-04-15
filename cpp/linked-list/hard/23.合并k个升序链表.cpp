/*
 * @lc app=leetcode.cn id=23 lang=cpp
 *
 * [23] 合并K个升序链表
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
class Solution {
public:
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        ListNode dummy;
        auto cmp = [](ListNode* l, ListNode* r){return l->val > r->val;};
        std::priority_queue<ListNode*, std::vector<ListNode*>, decltype(cmp)> q(cmp);
        // 每个链表的头节点加入优先队列
        for (auto e : lists) {
            if (e != nullptr) {
                q.push(e);
            }
        }

        auto p = &dummy;
        while (!q.empty()) {
            p->next = q.top();
            p = p->next;
            q.pop();
            if (p->next) {
                q.push(p->next);
            }
        }
        return dummy.next;
    }
    /*
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        if (lists.empty()) {
            return nullptr;
        }

        if (lists.size() < 2) {
            return lists[0];
        }
        
        ListNode* result = nullptr;
        for (auto e : lists) {
            result = merge2Lists(e, result);
        }

        return result;
    }

    ListNode* merge2Lists(ListNode* list1, ListNode* list2) {
        ListNode dummy;
        auto p = &dummy;
        auto p1 = list1;
        auto p2 = list2;
        while (p1 && p2) {
            if (p1->val < p2->val) {
                p->next = p1;
                p1 = p1->next;
            } else {
                p->next = p2;
                p2 = p2->next;
            }
            p = p->next;
        }

        if (p1) {
            p->next = p1;
        }

        if (p2) {
            p->next = p2;
        }

        return dummy.next;
    }
    */
};
// @lc code=end

