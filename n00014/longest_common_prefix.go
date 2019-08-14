package n00014

/**
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 * 如果不存在公共前缀，返回空字符串 ""。
 * 找出字符串数组中长度最小的字符串，此字符串用来当作标准对比的字符串。
 * 用两个循环匹配最长前缀
 * 外循环，从i=0到 最小字符串 的长度
 * 内循环从j=0到字符串数组长度，负责将字符串数组的每一个字符串从0截取到i长度与minStr对比
 *
 */


func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}else if len(strs) == 1 {
		return strs[0]
	}

	for i:=0;i<len(strs[0]);i++ {

		for j:=1;j<len(strs);j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]

}
