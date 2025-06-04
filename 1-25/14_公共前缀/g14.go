package main

import (
	"fmt"
	"math"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	isCommonPrefix := func(length int) bool {
		str := strs[0][:length]
		for i := 1; i < len(strs); i++ {
			for j := 0; j < length; j++ {
				if strs[i][j] != str[j] {
					return false
				}
			}
		}
		return true
	}
	mid := 0
	minLen := math.MaxInt
	for _, str := range strs {
		t := len(str)
		if t < minLen {
			minLen = t
		}
	}
	low, heigh := 0, minLen
	for low <= heigh {
		mid = (low + heigh) / 2
		if isCommonPrefix(mid) {
			low = mid + 1
		} else {
			heigh = mid - 1
		}
	}
	if low == 0 {
		return ""
	}
	return strs[0][:low-1]
}
func main() {
	fmt.Println(longestCommonPrefix([]string{"a"}))
}
