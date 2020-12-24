package n00135

func candy(ratings []int) int {

	var ratingsLen = len(ratings)
	// 左遍历 额外空间
	left := make([]int, ratingsLen)

	// 右遍历 额外空间
	right := make([]int, ratingsLen)

	// 初始值
	for i := 0; i < ratingsLen; i++ {
		left[i] = 1
		right[i] = 1
	}

	for i := 1; i < ratingsLen; i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		}
	}
	count := left[ratingsLen-1]
	for i := ratingsLen - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		}
		if left[i] >= right[i] {
			count += left[i]
		} else {
			count += right[i]
		}
	}

	return count
}
