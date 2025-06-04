package main

import "fmt"

func convert(s string, numRows int) string {
	s1 := make([][]int32, numRows)
	flag := -1
	i := 0
	for i, _ := range s1 {
		s1[i] = make([]int32, 0)
	}
	for _, v := range s {
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		s1[i] = append(s1[i], v)
		if numRows > 1 {
			i += flag

		}
	}
	res := make([]int32, 0)
	for _, v := range s1 {

		res = append(res, v...)
	}

	//fmt.Printf("\n")
	return string(res)
}
func main() {
	fmt.Printf("%s", convert("leetcod", 3))
}
