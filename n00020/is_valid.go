package n00020

func isValid(s string)bool{
	if len(s) == 0 {
		return true
	}

	if len(s)%2 == 1 {
		return false
	}

	m := make(map[byte]byte)
	m['{'] = '}'
	m['['] = ']'
	m['('] = ')'

	temStr := ""

	for i,_ := range s {
		if temStr == "" {
			temStr += s[i:i+1]
			continue
		}

		if _,ok := m[s[i]];!ok {
			if m[temStr[len(temStr) -1]] != s[i] {
				return false
			}
			temStr = temStr[:len(temStr) - 1]
		}else {
			temStr = temStr + s[i:i +1]
		}
	}

	return temStr == ""

}
