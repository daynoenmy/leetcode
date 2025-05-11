package main

import "fmt"

func maxArea(height []int) int {
	maxs := 0
	for l, r := 0, len(height)-1; l < len(height); {

		if l >= r {
			break
		}
		area := min(height[r], height[l]) * (r - l)
		if area > maxs {
			maxs = area
		}
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return maxs
}
func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
