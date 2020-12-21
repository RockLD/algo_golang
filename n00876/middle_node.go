package n00876

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	length := 0

	// 求长度
	for cur != nil {
		length++
		cur = cur.Next
	}
	step := length/2 + 1
	// 找到节点
	for i := 1; i < step; i++ {
		head = head.Next
	}
	return head
}
