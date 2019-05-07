package n00024
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 type ListNode struct {
	 Val int
	 Next *ListNode
 }
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var pre *ListNode
	cur := head
	head = cur.Next
	for ;cur != nil && cur.Next != nil ;cur = cur.Next {
		// next := cur.Next
		if pre != nil {
			pre.Next = cur.Next
		}
		cur.Next,cur.Next.Next,pre = cur.Next.Next,cur,cur
	}

	return head
}