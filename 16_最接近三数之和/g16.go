package main

import (
	"fmt"
	"math"
	"sort"
)

func twoSum(nums []int, target, loc, start, end int, res *int) {
	v, mv := 0, int(math.Abs(float64(*res-target)))
	for start < end {
		t := nums[loc] + nums[start] + nums[end] - target
		v = int(math.Abs(float64(t)))
		if t == 0 {
			*res = nums[loc] + nums[start] + nums[end]
			return
		} else if t > 0 {
			if v < mv {
				mv = v
				*res = nums[loc] + nums[start] + nums[end]
			}
			end--
		} else {
			if v < mv {
				mv = v
				*res = nums[loc] + nums[start] + nums[end]
			}
			start++
		}
	}

}
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		twoSum(nums, target, i, i+1, len(nums)-1, &res)
		//fmt.Println(res)

	}
	return res
}
func main() {
	fmt.Println(threeSumClosest([]int{0, 1, 2}, 3))
}
