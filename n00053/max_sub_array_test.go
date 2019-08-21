package n00053

import "testing"

func TestMaxSubArray(t *testing.T) {
	input := []int{-2,1,-3,4,-1,2,1,-5,4}

	res := maxSubArray(input)

	if res != 6 {
		t.Error("err")
	}
}
