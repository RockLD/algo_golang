package n00621

import "testing"

func TestLeastInterval(t *testing.T) {
	task := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	res := leastInterval(task, 2)
	t.Fatal(res)
}
