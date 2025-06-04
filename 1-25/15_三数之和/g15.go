package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target, start, end int, res *[][]int) {

	for start < end {
		if target+nums[start]+nums[end] == 0 {
			*res = append(*res, []int{target, nums[start], nums[end]})
			for start < end && nums[start] == nums[start+1] {
				start++
			}
			start++
			for start < end && nums[end] == nums[end-1] {
				end--
			}
			end--
		} else if target+nums[start]+nums[end] > 0 {
			end--
		} else if target+nums[start]+nums[end] < 0 {
			start++
		}
	}

}
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	s := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		twoSum(nums, nums[i], i+1, len(nums)-1, &s)
	}
	return s
}
func main() {

	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
