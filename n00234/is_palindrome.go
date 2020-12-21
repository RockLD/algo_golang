package n00234

/**
 * Definition for singly-linked list.
 * 请判断一个链表是否为回文链表
 * 思路
 * 遍历一遍链表,得到链表长度n,根据长度的奇偶,找到中间节点,将左半边的链表反转,然后从中间节点分两个方向向左右两边遍历,是否是回文;对左半部分链表进行反转,还原为最初的链表
 * 只需要固定的若干个临时变量,不需要额外开辟空间
 * 时间复杂度为O(n),空间复杂度为O(1)
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	isPalindrome := true

	length := 0

	for cur := head; cur != nil; cur = cur.Next {
		length++
	}

	// 前半部分反转
	step := length / 2
	var pre *ListNode
	cur := head

	for i := 1; i <= step; i++ {
		cur.Next, pre, cur = pre, cur, cur.Next
	}

	mid := cur

	var left, right *ListNode = pre, nil
	if length%2 == 0 { // 长度为偶数
		right = mid
	} else {
		right = mid.Next
	}

	// 从中间向两侧遍历对比
	for left != nil && right != nil {
		if left.Val != right.Val {
			isPalindrome = false
			break
		}
		left = left.Next
		right = right.Next
	}

	cur = pre
	pre = mid
	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	return isPalindrome
}
