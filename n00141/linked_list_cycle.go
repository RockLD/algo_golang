package n00141

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

func hasCycle(head *ListNode) bool {
	fast , slow := head,head
	for fast != nil && slow != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}