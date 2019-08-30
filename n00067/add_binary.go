package n00067

func addBinary(a, b string) string {
	var res string
	//进位标志
	carry := false

	aLen := len(a)

	bLen := len(b)

	//用 0 把两个字符串长度补齐
	if aLen > bLen {
		dif := aLen - bLen

		for start := 1; start <= dif; start++ {
			b = "0" + b
		}
	} else {
		dif := bLen - aLen

		for start := 1; start <= dif; start++ {
			a = "0" + a
		}

	}

	//循环运算
	for i := (len(a) - 1); i >= 0; i-- {
		//0-0		& = 0
		//0 -1      & = 0
		//1 - 0	    & = 0
		//1 -1      & = 1
		if a[i] == '0' && b[i] == '0' {

			println("a")
			if carry {
				res = "1" + res
				carry = false
			} else {
				res = "0" + res
			}
		} else if a[i] == '1' && b[i] == '1' {
			println("b")
			if carry {
				res = "1" + res
			} else {
				carry = true
				res = "0" + res
			}
		} else {
			println("c")
			if carry {
				res = "0" + res
			} else {
				res = "1" + res
			}
		}
	}

	//最后如果有进位，则前面加 1
	if carry {
		println("d")
		res = "1" + res
		carry = false
	}
	println(res)

	return res
}
