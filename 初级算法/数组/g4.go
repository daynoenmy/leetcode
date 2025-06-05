package main

import "fmt"

func containsDuplicate(nums []int) bool {
	flags := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if flags[nums[i]] == false {
			flags[nums[i]] = true
		} else {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(containsDuplicate([]int{1, 1, 3, 4, 5}))
}
