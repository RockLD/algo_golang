package n00028


func strStr(haystack,needle string) int {
	if len(needle) == 0 {
		return 0
	}

	if len(haystack) < len(needle) {
		return -1
	}

	if haystack == needle {
		 return 0
	}

	for i:=0;i<=(len(haystack)-len(needle));i++ {
		if haystack[i:len(needle)+i] == needle {
			return i
		}
	}

	return -1
}
