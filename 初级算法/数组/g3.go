package main

func reverseArray(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
func rotate(nums []int, k int) {
	if len(nums) < k {
		k = k % len(nums)
	}
	reverseArray(nums[0 : len(nums)-k])
	//fmt.Println(nums)
	reverseArray(nums[len(nums)-k : len(nums)])
	//fmt.Println(nums)
	reverseArray(nums)
	//fmt.Println(nums)
}
func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}
