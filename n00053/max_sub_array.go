package n00053

/**最大子序和----动态规划算法**/


func maxSubArray(nums []int) int {
	max := int(^(int(^uint(0) >> 1)))

	tmp := 0
	for _, v := range nums {
		if tmp > 0 {
			tmp += v
		}else {
			tmp = v
		}
		if max < tmp {
			max = tmp
		}
	}

	return max

}
