package main

import "fmt"

func removeElement(nums []int, val int) int {
	n := len(nums)
	l, r := 0, 0
	for r < n {
		if nums[r] != val {
			nums[l] = nums[r]
			l++
		}
		r++
	}
	return l
}
func main() {
	arr := []int{3, 2, 2, 3}
	fmt.Println(removeElement(arr, 3))
	fmt.Println(arr)
}
