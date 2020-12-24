package n00143

/**
 * 重排链表
 * Definition for singly-linked list.
 * 找中间节点，断开链表后，将后半部分链表反转，最后将两个链表合并
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	nodes := []*ListNode{}

	cur := head

	for cur != nil {
		nodes = append(nodes, cur)
		cur = cur.Next
	}

	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil

}
