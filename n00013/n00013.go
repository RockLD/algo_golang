package n00013
/**
	罗马数字转整数
	字符          数值
	I             1
	V             5
	X             10
	L             50
	C             100
	D             500
	M             1000
 */

 func RomanToInt(s string)int{
 	tagVal := map[string]int{"I":1,"V":5,"X":10,"L":50,"C":100,"D":500,"M":1000}
	var num int
	for i:=0;i<len(s);i++{
		if (i+1 >= len(s)) || (tagVal[string(s[i+1])] <= tagVal[string(s[i+1])]) {
			num += tagVal[string(s[i])]
		}else{
			num -= tagVal[string(s[i])]
		}
	}
 	return num
 }