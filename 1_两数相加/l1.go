package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		k, v := m[target-nums[i]]
		if v {
			return []int{i, k}
		}
		m[nums[i]] = i
	}
	return []int{0, 0}
}
func main() {

}
