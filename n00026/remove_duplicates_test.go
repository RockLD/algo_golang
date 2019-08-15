package n00026

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	arr := []int{1,1,2}
	result := removeDuplicates(arr)

	if result != 2 {
		t.Errorf("error,the result is %v",result)
	}

	arr = []int{0,0,1,1,1,2,2,3,3,4}
	result = removeDuplicates(arr)

	if result != 5 {
		t.Errorf("error, the result is %v",result)
	}

	arr = []int{1,2,2}
	result = removeDuplicates(arr)

	if result != 2 {
		t.Errorf("error,the result is %v",result)
	}
}