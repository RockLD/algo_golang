package n00203

/**
 * Definition for singly-linked list.
 * 需要构建一个新的链表，避免第一位删除的问题
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	newHead := &ListNode{-1, head}
	pre := newHead

	for pre.Next != nil {
		if pre.Next.Val == val {
			pre.Next = pre.Next.Next
		} else {
			pre = pre.Next
		}
	}
	return newHead.Next
}
