package main

import "fmt"

func plusOne(digits []int) []int {
	nums := len(digits)
	digits[nums-1] += 1
	n := make([]int, nums+1)
	for i := nums - 1; i >= 0; i-- {
		if digits[i] > 9 {
			digits[i] = 0
			if i != 0 {
				digits[i-1] += 1
			} else {
				digits[i] += 1
				n[0] = 1
				n[1] = 0
				return n
			}
		}
		n[i+1] = digits[i]
	}
	return n[1:]
}
func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
}
