package n00442

import "math"

func FindDuplicates(nums []int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		loc := int(math.Abs(float64(nums[i])))
		if nums[loc-1] > 0 {
			nums[loc-1] *= -1
		} else {
			result = append(result, loc)
		}
	}

	return result
}
