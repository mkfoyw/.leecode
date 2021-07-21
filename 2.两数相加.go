/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r := new(ListNode)
	r.Next = nil

	cur := r

	carry := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		n := new(ListNode)
		n.Val = carry % 10
		n.Next = nil
		cur.Next = n
		cur = n
		carry /= 10
	}

	if carry > 0 {
		n := new(ListNode)
		n.Val = carry
		n.Next = nil
		cur.Next = n
		cur = n
	}
	return r.Next
}

