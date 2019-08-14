package n00021

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	l1 := &ListNode{1,&ListNode{2,&ListNode{4,nil}}}
	l2 := &ListNode{1,&ListNode{3,&ListNode{4,nil}}}

	r := mergeTwoLists(l1,l2)
	var res []int
	for r != nil {
		res = append(res,r.Val)
		r = r.Next
	}
	expect := []int{1,1,2,3,4,4}

	if !reflect.DeepEqual(res,expect) {
		t.Errorf("result is err,the value is %v",res)
	}
}
