package n00389

func findTheDifference(s string, t string) byte {
	if s == "" {
		return []byte(t)[0]
	}

	m := make(map[int32]int, 256)

	var i int32
	for i = 0; i < 256; i++ {
		m[i] = 0
	}

	for _, v := range s {
		m[v]++
	}

	for _, v := range t {
		m[v]--
		if m[v] < 0 {
			return byte(v)
		}
	}
	return 0
}
