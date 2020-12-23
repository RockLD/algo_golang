package n01669

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	// 从第一个链表开始
	cur := list1

	// 找第一个断点
	for i := 0; i < (a - 1); i++ {
		cur = cur.Next
	}

	// 找第二个断点
	cur1 := cur.Next
	for i := 0; i < (b - a + 1); i++ {
		cur1 = cur1.Next
	}

	// 在第一个断点出，连接第二个链表，从此处开始，cur属于第二个链表的节点
	cur.Next = list2
	for cur.Next != nil {
		cur = cur.Next
	}
	// 第二个链表尾接点，连接第一个链表的第二个断点
	cur.Next = cur1
	return list1
}
