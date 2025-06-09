package main

import "fmt"

func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i-j] = nums[i]
		} else {
			j++
		}
	}
	for i := len(nums) - 1; j > 0; j-- {
		nums[i] = 0
		i--
	}
}

func main() {
	arr := []int{0, 1, 0, 3, 4, 5}
	moveZeroes(arr)
	fmt.Println(arr)
}
