package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	var left, right []int

	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	// 递归排序左右子数组
	left = quickSort(left)
	right = quickSort(right)

	// 合并结果
	return append(append(left, pivot), right...)
}

func intersectMy(nums1 []int, nums2 []int) []int {
	nums1 = quickSort(nums1)
	nums2 = quickSort(nums2)
	nums := make([]int, 0)
	i, j := 0, 0
	fmt.Println(nums1, nums2)
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			nums = append(nums, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		}
	}
	return nums
}
func intersect(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	nums := make([]int, 0)
	for _, v := range nums1 {
		m[v]++
	}
	for _, v := range nums2 {
		if m[v] > 0 {
			nums = append(nums, v)
			m[v]--
		}

	}
	return nums

}

func main() {

	nums := intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})
	fmt.Println(nums)
}
