package n00035


func searchInsert(nums []int, target int)int{
	left := 0
	right := len(nums)

	middle := 0

	for left < right && left < len(nums) {
		middle = (left + right) >> 1
		if target > nums[middle] {
			left = middle + 1
		}else if target < nums[middle] {
			right = middle
		}else {
			return middle
		}
	}
	return left
}