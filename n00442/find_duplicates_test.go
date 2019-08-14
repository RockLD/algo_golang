package n00442

import (
	"reflect"
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	var input = []int{4,3,2,7,8,2,3,1}
	res := FindDuplicates(input)
	expect := []int{2,3}
	if !reflect.DeepEqual(res,expect) {
		t.Errorf("The result is err,the result is %v",res)
	}
}