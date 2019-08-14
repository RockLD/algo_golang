package n00020

import "testing"

func TestIsValid(t *testing.T) {
	var res bool
	r := "()"
	res = isValid(r)
	if res != true {
		t.Error("failed")
	}
	r1 := "()[]{}"
	res = isValid(r1)
	if res != true {
		t.Error("failed")
	}
	r2 := "(]"
	res = isValid(r2)
	if res != false {
		t.Error("failed")
	}
	r3 := "([)]"
	res = isValid(r3)
	if res != false {
		t.Error("failed")
	}
	r4 := "{[]}"
	res = isValid(r4)
	if res != true {
		t.Error("failed")
	}
}
