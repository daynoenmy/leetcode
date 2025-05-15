package main

import "fmt"

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func tree(digits string, index int, combination string, res *[]string) {
	if index == len(digits) {
		*res = append(*res, combination)
	} else {
		digit := string(digits[index])
		letterMap := phoneMap[digit]
		letterCount := len(letterMap)
		for i := 0; i < letterCount; i++ {
			tree(digits, index+1, combination+string(letterMap[i]), res)
		}
	}
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	res := make([]string, 0)
	tree(digits, 0, "", &res)
	return res
}

func main() {

	fmt.Println(letterCombinations("22"))
}
