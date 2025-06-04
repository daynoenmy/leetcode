package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	res := 0
	f := 1
	fnum := false
	for i := 0; i < len(s); i++ {

		if s[i] == ' ' {
			if fnum {
				return res * f
			}
			continue
		}
		if s[i] == '-' {
			if fnum {
				return res * f
			}
			f = -1
			fnum = true
			continue
		}
		if s[i] == '+' {
			if fnum {
				return res * f
			}
			f = 1
			fnum = true
			continue
		}
		if '0' <= s[i] && s[i] <= '9' {
			fnum = true
			if res*f < math.MinInt32/10 {
				return math.MinInt32
			}
			if res*f > math.MaxInt32/10 {
				return math.MaxInt32
			}
			if res*f == math.MaxInt32/10 && s[i]-'0' > 7 {
				return math.MaxInt32
			}
			if res*f == math.MinInt32/10 && s[i]-'0' > 8 {
				return math.MinInt32
			}
			res = res*10 + int(s[i]-'0')
		} else {
			return res * f
		}
	}
	return res * f
}
func main() {
	//"2147483647"
	fmt.Println(myAtoi("2147483646"))
	fmt.Println(math.MinInt32)
}
