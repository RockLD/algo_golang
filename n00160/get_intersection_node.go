package n00160

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	cur1, cur2 := headA, headB
	for cur1 != cur2 {
		if cur1 != nil {
			cur1 = cur1.Next
		} else {
			cur1 = headA
		}
		if cur2 != nil {
			cur2 = cur2.Next
		} else {
			cur2 = headB
		}
	}
	return cur1
}
