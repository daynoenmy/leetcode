package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	fast, slow := 1, 1
	for fast < len(nums) {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
func main() {
	arr := []int{1, 1, 1, 2, 2, 3, 1}
	fmt.Println(removeDuplicates(arr))
	fmt.Println(arr)
}
