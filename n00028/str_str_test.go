package n00028

import "testing"

func TestStrStr(t *testing.T) {
	haystack := "a"
	needle := "a"
	result := strStr(haystack, needle)
	if result != 0 {
		t.Errorf("err,is %v",result)
	}
}