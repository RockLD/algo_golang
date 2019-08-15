package n00027

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}

	return i
}


//给定 nums = [3,2,2,3], val = 3,
//给定 nums = [0,1,2,2,3,0,4,2], val = 2,