package n00387

/**
 * 将每个字符与byte类型的'a'的差值作为索引
 */
func firstUniqChar(s string) int {

	if len(s) == 0 {
		return -1
	}
	if len(s) == 1 {
		return 0
	}
	res := make([]int, 26)

	for _, val := range s {
		res[val-'a']++
	}

	for i, v := range s {
		if res[v-'a'] == 1 {
			return i
		}
	}
	return -1

}
