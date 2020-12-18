package n00083

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 * 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次
 * 输入 1->1->2
 * 输出 1->2
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p := head
	// 循环链表
	for p.Next != nil {
		// 如果后面一个节点的值和前面相等，就删除
		if p.Val == p.Next.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}

	return head
}
