package n00083

import (
	"fmt"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	l := makeList()

	r := deleteDuplicates(l)
	r1 := r
	for r1.Next != nil {
		fmt.Println("val=", r1.Val)
		r1 = r1.Next
	}
	fmt.Println("val=", r1.Val)

}

func makeList() *ListNode {
	var num = []int{1, 1, 2, 3}
	res := &ListNode{Val: num[0]}

	temp := res

	for i := 1; i < len(num); i++ {
		temp.Next = &ListNode{Val: num[i]}
		temp = temp.Next
	}
	return res
}
