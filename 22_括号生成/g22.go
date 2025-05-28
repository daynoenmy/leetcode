package main

import "fmt"

func f(l, r int, str string, res *[]string) {
	if l == 0 && r == 0 {
		*res = append(*res, str)
		return
	}
	if l > 0 {
		str += "("
		f(l-1, r, str, res)
		str = str[:len(str)-1]
	}
	if r > l {
		str += ")"
		f(l, r-1, str, res)
		str = str[:len(str)-1]
	}

}
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	f(n, n, "", &res)
	return res
}
func main() {
	fmt.Println(generateParenthesis(2))
}
