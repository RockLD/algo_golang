package n00014

import (
	"testing"
)

func TestLongCommonPrefix(t *testing.T) {
	input1 := []string{"flower","flow","flight"}
	r := longestCommonPrefix(input1)
	if r != "fl" {
		t.Errorf("error of input1 is :%#v\n",r)
	}

	input2 := []string{"dog","racecar","car"}
	r2 := longestCommonPrefix(input2)
	if r2 != "" {
		t.Errorf("error of input2 is :%#v\n",r2)
	}

	t.Log("success")

}