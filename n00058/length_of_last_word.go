package n00058

import "strings"

/*最后一个单词的长度*/


func lengthOfLastWord(s string) int {
	if s = strings.TrimSpace(s);len(s) == 0 {
		return 0
	}

	ans := 0
	bs := []byte(s)
	for i:=len(s)-1;i>=0;i-- {
		if bs[i] == ' ' {
			break
		}
		ans++
	}

	return ans

}

