package n00038

import "testing"

func TestCountAndSay(t *testing.T) {
	n := 1
	res := countAndSay(n)
	if res != "1" {
		t.Error("err")
	}
}

