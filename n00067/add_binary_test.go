package n00067

import "testing"

func TestAddBinary(t *testing.T) {
	a := "1"
	b := "111"
	res := addBinary(a, b)
	t.Error(res)
}
