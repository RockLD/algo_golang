package n00933

type RecentCounter struct {
	data []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.data = append(this.data, t)
	start := 0
	for index, item := range this.data {
		if item >= t-3000 {
			start = index
			break
		}
	}
	this.data = this.data[start:]
	return len(this.data)
}
