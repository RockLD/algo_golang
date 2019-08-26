package n00066

func plusOne(digits []int)[]int {
	if digits[len(digits)-1] != 9 {
		digits[len(digits)-1]++
		return digits
	}
	i := len(digits) - 1
	for ;i>=0&&digits[i] == 9;i-- {
		digits[i] = 0
	}

	if i >= 0 {
		digits[i]++
	}else{
		digits = append([]int{1},digits...)
	}

	return digits
}
