package n00035

import "testing"

func TestSearchInsert(t *testing.T){
	nums := []int{1,3,5,6}
	target := 7
	res := searchInsert(nums,target)
	t.Error(res)
}