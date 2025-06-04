package main

import (
	"fmt"
	"math"
)

func reverse(x int) (int, bool) {
	res := 0
	for x > 0 {
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0, false
		}
		digit := x % 10
		x /= 10
		res = res*10 + digit
	}
	return res, true
}
func isPalindrome(x int) bool {
	r, f := reverse(x)
	if !f {
		return false
	}
	return r == x
}
func main() {
	fmt.Println(isPalindrome(12321))
}
