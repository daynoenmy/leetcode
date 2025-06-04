package main

import "fmt"

func romanToInt(s string) int {
	num := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 'I' {
			if i < len(s)-1 && (s[i+1] == 'V' || s[i+1] == 'X') {
				num -= 1
			} else {
				num += 1
			}
		} else if s[i] == 'V' {
			num += 5
		} else if s[i] == 'X' {
			if i < len(s)-1 && (s[i+1] == 'L' || s[i+1] == 'C') {
				num -= 10
			} else {
				num += 10
			}
		} else if s[i] == 'L' {
			num += 50
		} else if s[i] == 'C' {
			if i < len(s)-1 && (s[i+1] == 'D' || s[i+1] == 'M') {
				num -= 100
			} else {
				num += 100
			}
		} else if s[i] == 'M' {
			num += 1000
		} else if s[i] == 'D' {
			num += 500
		}
	}
	return num
}
func main() {
	fmt.Println(romanToInt("LVIII"))
}
