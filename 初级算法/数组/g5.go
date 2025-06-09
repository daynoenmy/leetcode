package main

func singleNumber(nums []int) int {
	a := 0
	for i := 0; i < len(nums); i++ {
		a ^= nums[i]
	}
	return a
}
func main() {

}
