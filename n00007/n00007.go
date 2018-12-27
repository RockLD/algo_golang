package n00007

import "math"

func reverse (x int) (r int) {


	for x != 0{
		r = r * 10 +  x % 10
		x = x / 10
		if r < math.MinInt32 || r > math.MaxInt32{
			return 0
		}
	}

	return r
}

