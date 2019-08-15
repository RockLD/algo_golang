package n00027

import "testing"

func TestRemoveElement(t *testing.T) {
	arr := []int{3,2,2,3}
	val := 3
	result := removeElement(arr,val)
	if result != 2 {
		t.Errorf("err")
	}
	arr = []int{0,1,2,2,3,0,4,2}
	val = 2
	result = removeElement(arr,val)
	if result != 5 {
		t.Errorf("err")
	}

}
