package n00038

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	ans := "11"

	for n > 2 {
		n--
		tmp := ""
		now := '+'
		count := 0
		for _,v := range  ans {
			if now == '+' {
				now = v
				count++
				continue
			}

			if now == v {
				count++
				continue
			}else {
				tmp += strconv.Itoa(count)
				tmp += string(now)
				count = 1
				now = v
				continue
			}
		}
		tmp += strconv.Itoa(count)
		tmp += string(now)
		ans = tmp
	}


	return ans

}
